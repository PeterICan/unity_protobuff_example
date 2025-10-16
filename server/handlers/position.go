package handlers

import (
	"fmt"
	"proto_buffer_example/server/third-party/antnet"
)

// PositionHandler will process player position updates
type PositionHandler struct{}

// HandlePositionUpdate processes position update messages from clients.
func (h *PositionHandler) HandlePositionUpdate(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	// For now, just log that we received a message.
	// The actual protobuf decoding will happen in the parser before this handler is called.
	fmt.Printf("Handling position update (Cmd: %d, Act: %d). Body size: %d\n", msg.Cmd(), msg.Act(), len(msg.Data))

	// Echo the message back to the client
	msgque.Send(msg)

	return true
}
