package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

//Args comment
type Args struct{}

//TimeServer comment
type TimeServer int64

// GiveServerTime comment
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	fmt.Println("Server is running on port 8000...")

	// Create a new RPC Server
	timeserver := new(TimeServer)

	//Register RPC Server
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	//Listen for requests on port 1234

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
