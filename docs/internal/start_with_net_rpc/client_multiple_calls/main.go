package main

import (
	"log"
	"net/rpc"
	"sync"
)

func hi(client *rpc.Client, message string)  {
	var reply string
	err := client.Call("echo.Message", message, &reply) ; if err != nil {
		log.Print(err)
	}
	log.Print(reply)
}

func main() {
	// 创建2个 tcp 链接，模拟多个机器像 rcp server 发送请求
	client1, err := rpc.Dial("tcp", "localhost:8888") ; if err != nil {
		panic(err)
	}
	defer client1.Close()
	client2, err := rpc.Dial("tcp", "localhost:8888") ; if err != nil {
		panic(err)
	}
	defer client2.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		hi(client1,"abc.")
		hi(client1,"xyz.")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		hi(client2,"123.")
		hi(client2,"fuck")
		wg.Done()
	}()
	wg.Wait()
}
