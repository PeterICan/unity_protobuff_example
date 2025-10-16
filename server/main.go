package main

import (
	"fmt"
	"proto_buffer_example/server/generated"
	"proto_buffer_example/server/handlers"
	"proto_buffer_example/server/third-party/antnet"
)

// Server defines the interface for a network server.
type Server interface {
	Start() error
}

// --- TCP Server Implementation ---

// tcpServer implements the Server interface for TCP connections.
type tcpServer struct {
	addr       string
	msgHandler *antnet.DefMsgHandler
	msgParser  *antnet.Parser
}

// NewTcpServer creates and configures a TCP server.
func NewTcpServer(addr string) Server {
	// 1. Initialize the Protobuf parser
	msgParser := &antnet.Parser{Type: antnet.ParserTypePB}
	msgParser.Register(
		byte(*generated.Cmd_CMD_POSITION.Enum()),
		byte(generated.ActPosition_ACT_POSITION_UPDATE),
		&generated.PlayerPosition{},
		nil,
	)

	// 2. Initialize the message handler
	msgHandler := &antnet.DefMsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	msgHandler.Register(
		byte(*generated.Cmd_CMD_POSITION.Enum()),
		byte(generated.ActPosition_ACT_POSITION_UPDATE),
		positionHandler.HandlePositionUpdate,
	)

	return &tcpServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  msgParser,
	}
}

func (s *tcpServer) Start() error {
	fmt.Printf("Starting antnet TCP server on %s with Protobuf parser\n", s.addr)
	return antnet.StartServer(s.addr, antnet.MsgTypeMsg, s.msgHandler, s.msgParser)
}

// --- WebSocket Server Implementation ---

// webSocketServer implements the Server interface for WebSocket connections.
type webSocketServer struct {
	addr       string
	msgHandler *antnet.DefMsgHandler
	msgParser  *antnet.Parser
}

// NewWebSocketServer creates and configures a WebSocket server.
func NewWebSocketServer(addr string) Server {
	// 1. Initialize the JSON parser
	jsonParser := &antnet.Parser{Type: antnet.ParserTypeJson}
	jsonParser.RegisterMsg(&generated.PlayerPosition{}, nil)

	// 2. Initialize the message handler
	msgHandler := &antnet.DefMsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	msgHandler.RegisterMsg(&generated.PlayerPosition{}, positionHandler.HandlePositionUpdate)

	return &webSocketServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  jsonParser,
	}
}

func (s *webSocketServer) Start() error {
	fmt.Printf("Starting antnet WebSocket server on %s with JSON parser\n", s.addr)
	return antnet.StartServer(s.addr, antnet.MsgTypeCmd, s.msgHandler, s.msgParser)
}

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
		server = NewWebSocketServer("ws://:7777")
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
