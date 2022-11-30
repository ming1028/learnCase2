package main

import (
	"github.com/learnCase2/rpc/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.RegisterName("MathService", new(server.MathService))
	if err != nil {
		log.Fatal("RegisterName error:", err)
	}
	// 基于http JSON RPC
	http.HandleFunc(rpc.DefaultRPCPath, func(writer http.ResponseWriter, request *http.Request) {
		conn, _, err := writer.(http.Hijacker).Hijack()
		if err != nil {
			log.Print("rpc hijacking ", request.RemoteAddr, ": ", err.Error())
			return
		}
		var connected = "200 Connected to JSON RPC"
		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")
		jsonrpc.ServeConn(conn)
	})
	// rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	// rpc.Accept(l)
	http.Serve(l, nil)
	// jsonrpc
	/*for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("jsonrpc.serve:accept:", err.Error())
			return
		}
		go jsonrpc.ServeConn(conn)
	}*/
}
