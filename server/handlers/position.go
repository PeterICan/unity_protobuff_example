package handlers

import (
	"encoding/json" // ADD THIS
	"fmt"
	"proto_buffer_example/server/generated"
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/third-party/antnet"
)

// PositionHandler will process player position updates
type PositionHandler struct{}

// HandlePositionMsgUpdate processes position update messages from clients.
func (h *PositionHandler) HandlePositionMsgUpdate(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	// For now, just log that we received a message.
	// The actual protobuf decoding will happen in the parser before this handler is called.
	fmt.Printf("Handling position update (Cmd: %d, Act: %d). Body size: %d\n", msg.Cmd(), msg.Act(), len(msg.Data))

	// Echo the message back to the client
	msgque.Send(msg)

	return true
}

func (h *PositionHandler) HandlePositionCmdUpdate(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	// For now, just log that we received a message.
	// The actual protobuf decoding will happen in the parser before this handler is called.
	fmt.Printf("Handling position update (Cmd: %d, Act: %d). Body size: %d\n", msg.Cmd(), msg.Act(), len(msg.Data))

	c2s := msg.C2S().(*generated.PlayerPosition)
	fmt.Printf("c2s:%v", c2s)

	// Echo the message back to the client
	msgque.Send(msg)

	return true

}

func (h *PositionHandler) HandleC2SPositionUpdate(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	c2sMsg, ok := msg.C2S().(*json_api.C2SPositionUpdate)
	if !ok {
		fmt.Printf("Error: Received message is not of type C2SPositionUpdate\n")
		return false
	}

	fmt.Printf("Received PositionUpdate request. Route: %s, X: %f, Y: %f, Z: %f\n", c2sMsg.Route, c2sMsg.X, c2sMsg.Y, c2sMsg.Z)

	// Create the correct response object
	response := &json_api.S2CPositionUpdate{
		Route:  c2sMsg.Route, // Echo back the route
		Status: "success",
		Error:  nil, // No error
	}

	// Manually marshal the response object to JSON bytes
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error marshalling response to JSON: %v\n", err)
		return false
	}

	// Create antnet.Message with the JSON data and no header
	antnetMsg := &antnet.Message{
		Data: jsonData,
		Head: nil, // Explicitly set Head to nil for WebSocket JSON
	}

	// Send the response back to the client
	msgque.Send(antnetMsg)

	return true
}
