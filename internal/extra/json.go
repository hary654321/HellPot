package extra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func WriteJson(path string, data interface{}) {

	// if global.ServerSetting.RunMode == "debug" {
	// 	prettyData, err1 := xjson.PrettyDumps(data)
	// 	if err1 != nil {
	// 		slog.Println(slog.DEBUG, err1)
	// 	}
	// 	slog.Println(slog.DEBUG, path, prettyData)
	// }

	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)

	err := enc.Encode(data)
	if err != nil {
		fmt.Printf("数据类型: %T, 数值: %v\n", err, err)
	}

	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
		fmt.Printf("数据类型: %T, 数值: %v\n", err, err)
	}

	//jsonBuf := append([]byte(result),[]byte("\r\n")...)
	f.Write(buf.Bytes())

}

func WriteJsonAny(path string, m map[string]interface{}) {

	currentTime := time.Now()
	// 将时间转换为毫秒时间戳
	milliseconds := currentTime.UnixNano() / int64(time.Millisecond)
	m["timestamp"] = milliseconds
	m["type"] = "connect"
	m["name"] = "HellPot"
	m["app"] = "HellPot"
	m["UUID"] = "<UUID>"

	WriteJson(path, m)
}
