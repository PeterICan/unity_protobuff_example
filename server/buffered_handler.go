package main

import (
	"fmt"
	"proto_buffer_example/server/handlers"
	"proto_buffer_example/server/third-party/antnet"
)

// BufferedMsgHandler embeds antnet.DefMsgHandler and overrides lifecycle methods
// to manage message buffers for each connection.
type BufferedMsgHandler struct {
	antnet.DefMsgHandler
}

// OnNewMsgQue is called when a new connection is established.
func (h *BufferedMsgHandler) OnNewMsgQue(msgque antnet.IMsgQue) bool {
	handlers.BufferMutex.Lock()
	defer handlers.BufferMutex.Unlock()

	connId := msgque.Id()
	handlers.ConnectionBuffers[connId] = &handlers.MessageBuffer{
		C2S_Queue: make([]*antnet.Message, 0),
		S2C_Queue: make([]*antnet.Message, 0),
	}

	fmt.Printf("Connection %d: Message buffer created.\n", connId)
	return true
}

// OnDelMsgQue is called when a connection is closed.
func (h *BufferedMsgHandler) OnDelMsgQue(msgque antnet.IMsgQue) {
	handlers.BufferMutex.Lock()
	defer handlers.BufferMutex.Unlock()

	connId := msgque.Id()
	delete(handlers.ConnectionBuffers, connId)

	fmt.Printf("Connection %d: Message buffer cleared.\n", connId)
}
