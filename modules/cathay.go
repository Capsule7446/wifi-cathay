package modules

import (
	"net/http"
	"strings"
	"time"
)

var client *http.Client
var isFlag = false

func init() {
	client = &http.Client{
		Timeout: time.Second,
	}
}

func Login() {
	// 无网状态
	if !Ping(2 * time.Second) {
		// 记录断网
		go NoHasNetwork()
		isFlag = false
		client.Post(ConfigData.Url, "application/json", strings.NewReader(""))
	} else { // 有网状态
		// 上一次无网络
		if isFlag == false {
			// 记录连接成功
			go HasNetwork()
			isFlag = true
		}
	}
}
