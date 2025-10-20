package main

import (
	"fmt"
	"proto_buffer_example/server/generated"
	"proto_buffer_example/server/handlers"
	"proto_buffer_example/server/third-party/antnet"
)

// --- WebSocket Server Implementation ---

// webSocketServer implements the Server interface for WebSocket connections.
type webSocketServer struct {
	addr       string
	msgHandler *handlers.MsgHandler
	msgParser  *antnet.Parser
}

// NewWebSocketServer creates and configures a WebSocket server.
func NewWebSocketServer(addr string) Server {
	// 1. Initialize the Protobuf parser
	pbParser := &antnet.Parser{}
	pbParser.Type = antnet.ParserTypeJson
	pbParser.RegisterMsg(&generated.PlayerPosition{}, nil)

	// 2. Initialize the message handler
	msgHandler := &handlers.MsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	msgHandler.RegisterMsg(&generated.PlayerPosition{}, positionHandler.HandlePositionCmdUpdate)

	return &webSocketServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  pbParser,
	}
}

func (s *webSocketServer) Start() error {
	fmt.Printf("Starting antnet WebSocket server on %s with Protobuf parser\n", s.addr)
	return antnet.StartServer(s.addr, antnet.MsgTypeCmd, s.msgHandler, s.msgParser)
}
