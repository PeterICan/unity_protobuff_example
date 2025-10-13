package handlers

import (
    "fmt"
    "github.com/magiclvzs/antnet"
    // We will need our generated protobuf types here later
    // "proto_buffer_example/server/generated"
)

// LoginHandler will process login requests (Message ID 1)
type LoginHandler struct{}

func (h *LoginHandler) Handle(ctx *antnet.Context) {
    // When we have the proto definition, we'll deserialize here.
    // For now, just log the raw message.
    fmt.Printf("Handling login request (Message ID: %d). Body size: %d\n", ctx.Msg.ID(), len(ctx.Msg.Body()))

    // Create a response message (Message ID 2)
    // The body would be a serialized LoginResponse protobuf message
    response, err := ctx.Session().NewMessage(2, []byte("Login OK from handler")) // Placeholder body
    if err != nil {
        fmt.Printf("Failed to create response message: %v\n", err)
        return
    }
    // Send the response
    ctx.Reply(response)
}
