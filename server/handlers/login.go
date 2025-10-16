package handlers

import (
	"proto_buffer_example/server/third-party/antnet"
)

// LoginHandler will process login requests (Message ID 1)
type LoginHandler struct{}

func (h *LoginHandler) Handle(msgCtx *antnet.DefMsgHandler) {
	// TODO : Implement login logic here

}

func (h *LoginHandler) Register(msgHandler *antnet.DefMsgHandler) {
	// TODO : Register login handler here
}
