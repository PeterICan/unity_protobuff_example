package server

import (
	"fmt"
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/internal/player/container"
	"proto_buffer_example/server/internal/proto/handlers"
	"proto_buffer_example/server/third-party/antnet"
	customize "proto_buffer_example/server/tools/customize"
)

// --- WebSocket Server Implementation ---

// webSocketServer implements the Server interface for WebSocket connections.
type webSocketServer struct {
	addr       string
	msgHandler *customize.MsgHandler
	msgParser  *antnet.Parser // Revert type to *antnet.Parser
	container.GamerContainer
}

// NewWebSocketServer creates and configures a WebSocket server.
func NewWebSocketServer(addr string) Server {
	// 1. Initialize the base antnet.Parser (which acts as the factory)
	baseParser := &antnet.Parser{}
	baseParser.Type = antnet.ParserTypeCustom          // Set its type to Custom
	baseParser.ErrType = antnet.ParseErrTypeSendRemind // Example error handling

	// Instantiate our custom parser
	jsonRouteParser := customize.NewJsonRouteParser(baseParser)

	// Set the custom parser using the new public setter
	baseParser.SetIParser(jsonRouteParser) // Use the new setter

	// Register C2S and S2C message types with our custom parser
	// Get the custom parser instance via the public Get() method
	customIParser := baseParser.Get()
	customJsonParser, ok := customIParser.(*customize.JsonRouteParser)
	if !ok {
		panic("Failed to cast IParser to *parser.JsonRouteParser")
	}
	customJsonParser.RegisterMsg("position/update", &json_api.C2SPositionUpdate{}, &json_api.S2CPositionUpdate{})
	customJsonParser.RegisterMsg("gamer_info/retrieve", &json_api.C2SGamerInfoRetrieve{}, &json_api.S2CGamerInfoRetrieve{})
	// 2. Initialize the message handler
	msgHandler := &customize.MsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	gamerInfoHandler := &handlers.GamerInfoHandler{}

	// Register handlers for the C2S message types
	msgHandler.RegisterMsg(&json_api.C2SPositionUpdate{}, positionHandler.HandleC2SPositionUpdate)
	msgHandler.RegisterMsg(&json_api.C2SGamerInfoRetrieve{}, gamerInfoHandler.HandleC2SGamerInfoRetrieve)

	return &webSocketServer{
		addr:       addr,
		msgHandler: msgHandler,
		msgParser:  baseParser, // Assign the baseParser (which is the factory)
	}
}

func (s *webSocketServer) StartServer() error {
	fmt.Printf("Starting antnet WebSocket server on %s with Custom JSON Route parser\n", s.addr)
	s.InitWebSocketServerBase(3345678, s.addr)
	return nil
	//return antnet.StartServer(s.addr, antnet.MsgTypeCmd, s.msgHandler, s.msgParser)
}
