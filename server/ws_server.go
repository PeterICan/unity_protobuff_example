package main

import (
	"fmt"
	"proto_buffer_example/server/generated/json_api"

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
	// pbParser.RegisterMsg(&generated.PlayerPosition{}, nil) // REMOVE THIS

	// ADD THESE
	pbParser.RegisterMsg(&json_api.C2SPositionUpdate{}, &json_api.S2CPositionUpdate{})
	pbParser.RegisterMsg(&json_api.C2SGamerInfoRetrieve{}, &json_api.S2CGamerInfoRetrieve{})

	// 2. Initialize the message handler
	msgHandler := &handlers.MsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	gamerInfoHandler := &handlers.GamerInfoHandler{} // ADD THIS

	// msgHandler.RegisterMsg(&generated.PlayerPosition{}, positionHandler.HandlePositionCmdUpdate) // REMOVE THIS

	// ADD THESE
	msgHandler.RegisterMsg(&json_api.C2SPositionUpdate{}, positionHandler.HandleC2SPositionUpdate)
	msgHandler.RegisterMsg(&json_api.C2SGamerInfoRetrieve{}, gamerInfoHandler.HandleC2SGamerInfoRetrieve)

	return &webSocketServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  pbParser,
	}
}

func (s *webSocketServer) Start() error {
	fmt.Printf("Starting antnet WebSocket server on %s with JSON parser\n", s.addr) // Updated log message
	return antnet.StartServer(s.addr, antnet.MsgTypeCmd, s.msgHandler, s.msgParser) // Changed MsgTypeMsg back to MsgTypeCmd
}
