package server

import (
	"fmt"

	"proto_buffer_example/server/generated"
	"proto_buffer_example/server/internal/proto/handlers"
	"proto_buffer_example/server/third-party/antnet"
)

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
		positionHandler.HandlePositionMsgUpdate,
	)

	return &tcpServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  msgParser,
	}
}

func (s *tcpServer) StartServer() error {
	fmt.Printf("Starting antnet TCP server on %s with Protobuf parser\n", s.addr)
	return antnet.StartServer(s.addr, antnet.MsgTypeMsg, s.msgHandler, s.msgParser)
}
