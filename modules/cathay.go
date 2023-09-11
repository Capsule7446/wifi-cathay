package modules

import (
	"log"
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
	log.Println(ConfigData.Url)
	if !Ping(2 * time.Second) {
		client.Post(ConfigData.Url, "application/json", strings.NewReader(""))
	}
}
