package handlers

import (
	"fmt"
	"proto_buffer_example/server/third-party/antnet"
)

type MsgHandler struct {
	antnet.DefMsgHandler
}

func (p *MsgHandler) OnNewMsgQue(msgque antnet.IMsgQue) bool {
	fmt.Printf("New message queue established: %v\n", msgque)
	return true
}

func (p *MsgHandler) OnDelMsgQue(msgque antnet.IMsgQue) {
	fmt.Printf("Message queue closed: %v\n", msgque)
}

func (p *MsgHandler) OnProcessMsg(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	fmt.Printf("Processing message (Cmd: %d, Act: %d). Body size: %d\n", msg.Cmd(), msg.Act(), len(msg.Data))
	return true
}

func (p *MsgHandler) OnConnectComplete(msgque antnet.IMsgQue, success bool) bool {
	fmt.Printf("Connection established: %v, Success: %v\n", msgque, success)
	return true
}

func (p *MsgHandler) GetHandlerFunc(msgque antnet.IMsgQue, msg *antnet.Message) antnet.HandlerFunc {
	fmt.Printf("Retrieving handler function for message queue: %v\n", msgque)
	return p.DefMsgHandler.GetHandlerFunc(msgque, msg)
}
