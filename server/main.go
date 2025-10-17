package main

import (
	"fmt"
	"proto_buffer_example/server/third-party/antnet"
)

func main() {
	// --- Configuration ---
	// Set serverType to "tcp" or "ws"
	serverType := "ws"
	// -------------------

	var server Server

	switch serverType {
	case "tcp":
		server = NewTcpServer("tcp://:6666")
	case "ws":
		server = NewWebSocketServer("ws://:7777/ws")
	default:
		fmt.Printf("Unknown server type: %s\n", serverType)
		return
	}

	// Start the selected server
	// The Start() method is blocking, so we run it in a goroutine
	// to allow WaitForSystemExit to catch the shutdown signal.
	go func() {
		if err := server.Start(); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
			panic(err)
		}
	}()

	// Wait for a signal to exit for graceful shutdown
	antnet.WaitForSystemExit()
}
