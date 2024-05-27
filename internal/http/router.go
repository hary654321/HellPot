package http

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/heffalump"
	"github.com/yunginnanet/HellPot/internal/config"
)

var log *zerolog.Logger

func getRealRemote(ctx *fasthttp.RequestCtx) string {
	xrealip := string(ctx.Request.Header.Peek(config.HeaderName))
	if len(xrealip) > 0 {
		return xrealip
	}
	return ctx.RemoteIP().String()
}

func hellPot(ctx *fasthttp.RequestCtx) {
	path, pok := ctx.UserValue("path").(string)
	if len(path) < 1 || !pok {
		path = "/"
	}
	localHost, localPort, _ := net.SplitHostPort(ctx.LocalAddr().String())
	remoteHost, remotePort, _ := net.SplitHostPort(ctx.RemoteAddr().String())
	currentTime := time.Now()
	// 将时间转换为毫秒时间戳
	milliseconds := currentTime.UnixNano() / int64(time.Millisecond)

	slog := log.With().
		Str("protocol", "http").
		Str("name", "HellPot").
		Str("app", "HellPot").
		Str("UUID", "<UUID>").
		Str("dest_ip", localHost).
		Str("src_ip", remoteHost).
		Int("dest_port", getStr2int(localPort)).
		Int("src_port", getStr2int(remotePort)).
		Int64("timestamp", milliseconds).
		Logger()

	for _, denied := range config.UseragentBlacklistMatchers {
		if strings.Contains(string(ctx.UserAgent()), denied) {
			ctx.Error("Not found", http.StatusNotFound)
			return
		}
	}

	extend := make(map[string]any)

	slog.Info().
		Interface("extend", extend).
		Str("type", "new-connect").Msg("NEW")

	s := time.Now()
	var n int64

	ctx.SetBodyStreamWriter(func(bw *bufio.Writer) {
		var err error
		var wn int64

		for {
			wn, err = heffalump.DefaultHeffalump.WriteHell(bw)
			n += wn
			if err != nil {
				slog.Trace().Err(err).
					Str("type", "error-connect").
					Msg("END_ON_ERR")
				break
			}
		}

		extend := make(map[string]any)
		extend["bytes"] = n
		extend["duration"] = time.Since(s)
		slog.Info().
			Interface("extend", extend).
			Str("type", "close-connect").
			Msg("close-connection")
	})

}

func getStr2int(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("转换失败:", err)
		return 0
	}

	fmt.Println("转换后的整数:", num)

	return num
}

func getSrv(r *router.Router) fasthttp.Server {
	if !config.RestrictConcurrency {
		config.MaxWorkers = fasthttp.DefaultConcurrency
	}

	log = config.GetLogger()

	return fasthttp.Server{
		// User defined server name
		// Likely not useful if behind a reverse proxy without additional configuration of the proxy server.
		Name: config.FakeServerName,

		/*
			from fasthttp docs: "By default request read timeout is unlimited."
			My thinking here is avoiding some sort of weird oversized GET query just in case.
		*/
		ReadTimeout:        5 * time.Second,
		MaxRequestBodySize: 1 * 1024 * 1024,

		// Help curb abuse of HellPot (we've always needed this badly)
		MaxConnsPerIP:      10,
		MaxRequestsPerConn: 2,
		Concurrency:        config.MaxWorkers,

		// only accept GET requests
		GetOnly: true,

		// we don't care if a request ends up being handled by a different handler (in fact it probably will)
		KeepHijackedConns: true,

		CloseOnShutdown: true,

		// No need to keepalive, our response is a sort of keep-alive ;)
		DisableKeepalive: true,

		Handler: r.Handler,
		Logger:  log,
	}
}

// Serve starts our HTTP server and request router
func Serve() error {
	log = config.GetLogger()
	l := config.HTTPBind + ":" + config.HTTPPort

	r := router.New()

	//配置不会走到这里
	if config.MakeRobots && !config.CatchAll {
		r.GET("/robots.txt", robotsTXT)
	}

	if !config.CatchAll {
		for _, p := range config.Paths {
			// log.Trace().Str("caller", "router").Msgf("Add route: %s", p)
			r.GET(fmt.Sprintf("/%s", p), hellPot)
		}
	} else {
		// log.Trace().Msg("Catch-All mode enabled...")
		r.GET("/{path:*}", hellPot)
	}

	srv := getSrv(r)

	//goland:noinspection GoBoolExpressions
	if !config.UseUnixSocket || runtime.GOOS == "windows" {
		// log.Info().Str("caller", l).Msg("Listening and serving HTTP...")
		return srv.ListenAndServe(l)
	}

	if len(config.UnixSocketPath) < 1 {
		log.Fatal().Msg("unix_socket_path configuration directive appears to be empty")
	}

	// log.Info().Str("caller", config.UnixSocketPath).Msg("Listening and serving HTTP...")
	return listenOnUnixSocket(config.UnixSocketPath, r)
}
