package main

import (
	"fmt" // Import the new package
	"proto_buffer_example/server/third-party/antnet"
)

func main() {

	addr := "tcp://:6666"
	msgHandler := &antnet.EchoMsgHandler{}
	// msgHandler := &customize.MsgHandler{}
	// msgParser := &antnet.Parser{Type: antnet.ParserTypePB}

	err := antnet.StartServer(addr, antnet.MsgTypeMsg, msgHandler, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		panic(err)
	}
	antnet.WaitForSystemExit()

	// // Register handlers from the handlers package
	// msgHandler.Register(1, &handlers.LoginHandler{})
	// msgHandler.Register(3, &handlers.PositionHandler{}) // Register the new handler

	// fmt.Println("Starting antnet server on tcp://:6666")
	// msgHandler.Run()
}
