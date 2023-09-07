package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/getlantern/systray"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

//go:embed embed/*
var embeddedFiles embed.FS

// Config 结构体用于映射 YAML 文件的内容
type Config struct {
	Time int    `yaml:"time"`
	Url  string `yaml:"url"`
}

var config Config

// $yyyy-MM-dd$/$HH-mm-ss$.png
func main() {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		return
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("无法解析 YAML 数据: %v", err)
	} else {
		systray.Run(onReady, onExit)
	}
}

func onReady() {
	systray.SetIcon(Image("bank"))
	systray.SetTitle("cathay")
	systray.SetTooltip("")
	start := systray.AddMenuItem("Start", "开始定时发送")
	start.SetIcon(Image("start"))
	stop := systray.AddMenuItem("Stop", "停止定时发送")
	stop.SetIcon(Image("stop"))
	exit := systray.AddMenuItem("Exit", "停止定时发送")
	exit.SetIcon(Image("exit"))
	go func(start, stop, exit *systray.MenuItem) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		client := &http.Client{Timeout: time.Second * 3}
		for {
			select {
			case <-start.ClickedCh:
				ctx, cancelFunc = context.WithCancel(context.Background())
				log.Println("start")
				if HasWifi("cathaybkguest") {
					ticker := time.NewTicker(time.Second * time.Duration(config.Time))
					go func(ticker *time.Ticker, ctx context.Context) {
						systray.SetIcon(Image("bankIsRun"))
						for {
							select {
							case <-ticker.C:
								if HasWifi("cathaybkguest") {
									Send(client)
									cancelFunc()
								}
							case <-ctx.Done():
								log.Println("ctx done")
								systray.SetIcon(Image("bank"))
								ticker.Stop()
								return
							}
						}
					}(ticker, ctx)
				}
			case <-stop.ClickedCh:
				log.Println("stop")
				systray.SetIcon(Image("bank"))
				cancelFunc()
			case <-exit.ClickedCh:
				log.Println("exit")
				systray.SetIcon(Image("bank"))
				cancelFunc()
				systray.Quit()
				return
			}
		}
	}(start, stop, exit)
}

func onExit() {
}

func HasWifi(str string) bool {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// 解析输出以获取WiFi名称
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "SSID") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				wifiName := strings.TrimSpace(parts[1])
				if strings.EqualFold(wifiName, str) {
					return true
				}
				break
			}
		}
	}
	return false
}
func Send(httpClient *http.Client) {
	defer func() {
		if e := recover(); e != nil {
		}
	}()
	log.Println("Send")
	httpClient.Post(config.Url, "application/json", strings.NewReader(""))
}
func Image(name string) []byte {
	fileContents, err := embeddedFiles.ReadFile(fmt.Sprintf("embed/%s.ico", name))
	if err != nil {
		fmt.Println("无法读取嵌入的图像文件:", err)
		return []byte{}
	}
	return fileContents
}
