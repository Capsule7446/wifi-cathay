package modules

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func HasNetwork() {
	Write("网络重新认证成功")
}

func NoHasNetwork() {
	Write("网络认证失效")
}
func Write(data string) {
	fileName := "LoginLog.txt"
	message := fmt.Sprintf("%v %s\n", time.Now().Format("2006-01-02 15:04:05"), data)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(message)
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}
