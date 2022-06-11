package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	todoService := NewTodoService()
	rpc.Register(todoService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
