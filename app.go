package main

import (
	"embed"
	"github.com/getlantern/systray"
	"log"
	"os"
	"time"
	c "wifi-cathay/modules"
)

//go:embed embed/*
var embedFiles embed.FS

func main() {
	// Read Config
	config, err := os.ReadFile("config.yml")
	if err != nil {
		return
	}
	c.Init(config)
	systray.Run(onReady, func() {})
}

func onReady() {
	systray.SetIcon(c.Image("bank", embedFiles))
	systray.SetTitle("cathay")
	systray.SetTooltip("")
	start := MenuItemStart("start", "开始定时发送")
	stop := MenuItemStart("stop", "关闭定时发送")
	exit := MenuItemStart("exit", "结束程式")
	go Run(start.ClickedCh, stop.ClickedCh, exit.ClickedCh)
}
func MenuItemStart(name string, tip string) *systray.MenuItem {
	start := systray.AddMenuItem(name, tip)
	start.SetIcon(c.Image(name, embedFiles))
	return start
}

func Run(start, end, exit chan struct{}) {
	var ticker *time.Ticker
	stopFunc := make(chan struct{})
	for c.IsCathayWIFI() {
		select {
		case <-start:
			log.Println("Start")
			if ticker == nil {
				log.Println("Start Success")
				ticker = time.NewTicker(time.Second * time.Duration(c.ConfigData.Time))
				go func(ticker *time.Ticker, stopFunc <-chan struct{}) {
					for ticker != nil {
						select {
						case <-ticker.C:
							log.Println("Ticker Run")
							go c.Login()
						}
					}
				}(ticker, stopFunc)
			}
		case <-end:
			if ticker != nil {
				ticker.Stop()
				ticker = nil
			}
			log.Println("End")
		case <-exit:
			if ticker != nil {
				ticker.Stop()
				ticker = nil
			}
			systray.Quit()
			log.Println("Exit")
		}
	}
}
