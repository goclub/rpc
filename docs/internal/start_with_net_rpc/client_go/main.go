package main

import (
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8888") ; if err != nil {
		panic(err)
	}
	defer client.Close()
	startTime := time.Now()
	var reply string
	// 通过 go.. 的两个 . 来模拟服务需要处理2秒
	messageCall := client.Go("echo.Message", "go..", &reply, nil)
	// 在此处的代码可以处理其他的逻辑，不会被 echo.Message 堵塞
	doSome()
	// 此处会因为 echo.Message 正在处理请求而堵塞
	result := <- messageCall.Done
	if result.Error != nil {
		log.Print(result)
		return
	}
	log.Print("Reply: " + reply)
	endTime := time.Now()
	log.Print("耗时：", endTime.Sub(startTime).String())
	// 这里的耗时会是 2s ，如果将 doSome() 中的 time.Second * 1 改成 time.Second * 3 则会变成耗时3秒
}

func doSome() {
	time.Sleep(time.Second * 1)
}

