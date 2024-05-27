package http

import (
	"net"

	"github.com/valyala/fasthttp"
	"github.com/yunginnanet/HellPot/internal/extra"
)

func robotsTXT(ctx *fasthttp.RequestCtx) {

	logmap := make(map[string]interface{})

	localHost, localPort, _ := net.SplitHostPort(ctx.LocalAddr().String())
	remoteHost, remotePort, _ := net.SplitHostPort(ctx.RemoteAddr().String())

	logmap["protocol"] = "tcp"
	logmap["src_ip"] = remoteHost
	logmap["src_port"] = remotePort
	logmap["dest_ip"] = localHost
	logmap["dest_port"] = localPort
	logmap["extend"] = ctx.Request

	extra.WriteJsonAny("log.json", logmap)

}
