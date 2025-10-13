package handlers

import (
    "fmt"
    "github.com/magiclvzs/antnet"
)

// PositionHandler will process player position updates (Message ID 3)
type PositionHandler struct{}

func (h *PositionHandler) Handle(ctx *antnet.Context) {
    // For now, just log the raw message.
    fmt.Printf("Handling position update (Message ID: %d). Body size: %d\n", ctx.Msg.ID(), len(ctx.Msg.Body()))

    // Position updates might not need a direct reply, or could be a simple ACK (Message ID 4)
    // No reply for now to simulate a fire-and-forget update.
}
