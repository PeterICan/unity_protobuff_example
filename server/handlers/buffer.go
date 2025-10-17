package handlers

import (
	"proto_buffer_example/server/third-party/antnet"
	"sync"
)

// MessageBuffer holds the C2S and S2C message queues for a single connection.
type MessageBuffer struct {
	C2S_Queue []*antnet.Message
	S2C_Queue []*antnet.Message
}

// ConnectionBuffers stores the buffers for all active connections.
// Key: connection ID (msgque.Id()), Value: pointer to the MessageBuffer.
var ConnectionBuffers = make(map[uint32]*MessageBuffer)

// BufferMutex protects concurrent access to the ConnectionBuffers map.
var BufferMutex = &sync.RWMutex{}
