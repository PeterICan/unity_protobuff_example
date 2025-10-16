package main

import (
	"fmt"
	"proto_buffer_example/server/generated"
	"proto_buffer_example/server/handlers"
	"proto_buffer_example/server/third-party/antnet"
)

func main() {
	addr := "tcp://:6666"

	// 1. Initialize the Protobuf parser
	msgParser := &antnet.Parser{Type: antnet.ParserTypePB}

	// 2. Register message types with the parser
	// Maps Cmd/Act to a specific Protobuf message type for automatic deserialization.
	msgParser.Register(
		byte(*generated.Cmd_CMD_POSITION.Enum()),
		byte(generated.ActPosition_ACT_POSITION_UPDATE),
		&generated.PlayerPosition{},
		nil,
	)

	// 3. Initialize the message handler
	msgHandler := &antnet.DefMsgHandler{}

	// 4. Register message handlers
	registerHandlers(msgHandler)

	// 5. Start the server
	err := antnet.StartServer(addr, antnet.MsgTypeMsg, msgHandler, msgParser)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		panic(err)
	}

	fmt.Println("Starting antnet server on tcp://:6666 with Protobuf parser")
	antnet.WaitForSystemExit()
}

// Placeholder for additional handler registrations if needed
func registerHandlers(msgHandler *antnet.DefMsgHandler) {
	positionHandler := &handlers.PositionHandler{}
	msgHandler.Register(
		byte(*generated.Cmd_CMD_POSITION.Enum()),
		byte(generated.ActPosition_ACT_POSITION_UPDATE),
		positionHandler.HandlePositionUpdate,
	)
}
