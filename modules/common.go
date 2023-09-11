package modules

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var WifiCommand *exec.Cmd

func init() {

}

func IsCathayWIFI() bool {
	WifiCommand = exec.Command("netsh", "wlan", "show", "interfaces")
	WifiCommand.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := WifiCommand.CombinedOutput()
	if err != nil {
		log.Println("Error:", err)
		return false
	}
	// 解析输出以获取WiFi名称
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if !strings.Contains(line, "SSID") {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		wifiName := strings.TrimSpace(parts[1])
		if strings.EqualFold(wifiName, ConfigData.Wifi) {
			return true
		}
	}
	return false
}
func Image(name string, embed embed.FS) []byte {
	fileContents, err := embed.ReadFile(fmt.Sprintf("embed/%s.ico", name))
	if err != nil {
		log.Println("无法读取嵌入的图像文件:", err)
		return []byte{}
	}
	return fileContents
}

func Ping(Timeout time.Duration) bool {
	http.DefaultClient.Timeout = Timeout
	_, err := http.Head("https://www.google.com/")
	return !(err != nil)
}
