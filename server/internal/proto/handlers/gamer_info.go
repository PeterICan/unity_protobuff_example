package handlers

import (
	"encoding/json" // ADD THIS
	"fmt"
	"proto_buffer_example/server/generated/json_api" // Import our new generate // Import our new generated types
	"proto_buffer_example/server/third-party/antnet"
)

// GamerInfoHandler will process gamer information requests
type GamerInfoHandler struct{}

// HandleC2SGamerInfoRetrieve processes C2SGamerInfoRetrieve messages from clients.
func (h *GamerInfoHandler) HandleC2SGamerInfoRetrieve(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	// The antnet parser should have already deserialized the message into msg.C2S()
	// We need to cast it to our specific C2SGamerInfoRetrieve type.
	c2sMsg, ok := msg.C2S().(*json_api.C2SGamerInfoRetrieve)
	if !ok {
		fmt.Printf("Error: Received message is not of type C2SGamerInfoRetrieve\n")
		return false
	}

	fmt.Printf("Received GamerInfoRetrieve request. Route: %s\n", c2sMsg.Route)

	// Create a dummy response for now
	response := &json_api.S2CGamerInfoRetrieve{
		Route:    c2sMsg.Route, // Echo back the route
		GamerId:  12345,
		NickName: "TestPlayer",
		Level:    10,
		Money:    1000,
		Gem:      500,
		Error:    nil, // No error
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
