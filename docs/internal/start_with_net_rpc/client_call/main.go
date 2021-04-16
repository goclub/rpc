package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8888") ; if err != nil {
		panic(err)
	}
	defer client.Close()
	var reply string
	err = client.Call("echo.Message", "nimoc", &reply) ; if err != nil {
		log.Print(err)
	}
	log.Print(reply)
}
