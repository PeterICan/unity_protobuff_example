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
	pbParser := &antnet.Parser{Type: antnet.ParserTypePB}
	// pbParser.Register(
	// 	uint8(*generated.Cmd_CMD_POSITION.Enum()),
	// 	uint8(generated.ActPosition_ACT_POSITION_UPDATE),
	// 	&generated.PlayerPosition{},
	// 	nil,
	// )
	pbParser.RegisterMsg(&generated.PlayerPosition{}, nil)

	// 2. Initialize the message handler
	msgHandler := &handlers.MsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	// msgHandler.Register(
	// 	uint8(*generated.Cmd_CMD_POSITION.Enum()),
	// 	uint8(generated.ActPosition_ACT_POSITION_UPDATE),
	// 	positionHandler.HandlePositionMsgUpdate,
	// )
	msgHandler.RegisterMsg(&generated.PlayerPosition{}, positionHandler.HandlePositionCmdUpdate)

	return &webSocketServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  pbParser,
	}
}

func (s *webSocketServer) Start() error {
	fmt.Printf("Starting antnet WebSocket server on %s with Protobuf parser\n", s.addr)
	fmt.Printf("Registering Cmd: %d, Act: %d\n", uint8(*generated.Cmd_CMD_POSITION.Enum()), uint8(generated.ActPosition_ACT_POSITION_UPDATE))
	// Note: antnet WebSocket server always uses MsgTypeCmd internally
	// The address format should be "ws://host:port/path" where /path is the WebSocket endpoint
	return antnet.StartServer(s.addr, antnet.MsgTypeMsg, s.msgHandler, s.msgParser)
}
