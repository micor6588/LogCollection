package main

import (
	"LogCollection/kafka"
	"LogCollection/taillog"
	"fmt"
	"time"
)

func main() {
	// 1.初始化kafka
	err := kafka.Init([]string{"127.0.0.1:29092"})
	if err != nil {
		fmt.Printf("init kafka failed ,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2.初始化taillog
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("init taillog failed, err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}
