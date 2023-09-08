package modules

import (
	"net/http"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second,
	}
}

func Login() {
	client.Post(ConfigData.Url, "application/json", strings.NewReader(""))
}
