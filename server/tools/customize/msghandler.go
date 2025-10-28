package customize

import (
	"log"
	"proto_buffer_example/server/internal/mediator"
	"proto_buffer_example/server/third-party/antnet"
)

type MsgStopFunc func(uint64) bool

/*
	封包處理模組 實現 antnet.IMsgHandler
*/

type MsgHandler struct {
	antnet.DefMsgHandler
	stopFunc MsgStopFunc
	//GamerContainer *container.GamerContainer
}

func (p *MsgHandler) SetStopFunc(callbackFunc MsgStopFunc) {
	//ctx := context.Background()
	//log.Debug(ctx, "SetStopFunc...")
	p.stopFunc = callbackFunc
}

func (p *MsgHandler) OnNewMsgQue(msgque antnet.IMsgQue) bool {
	//log.Info(ctx, "----- OnNewMsgQue GetUser:%d", msgque.GetUser())
	mediator.IGamerContainerModelMdr.NewGamerInitData(msgque)

	return true
}

func (p *MsgHandler) OnDelMsgQue(msgque antnet.IMsgQue) {
	//ctx := context.Background()
	if msgque.GetUser() == nil {
		return
	}

	gamerId := msgque.GetUser().(uint64)
	if gamerId == 0 {
		//log.Error(ctx, "----- OnDelMsgQue msgqueId:%d gamerId is zero", msgque.Id())
		return
	}
	log.Default().Println("玩家編號：", gamerId, "斷線 OnDelMsgQue called")
	mediator.IGamerContainerModelMdr.RemoveContainerGamer(gamerId)
	//event.MustFire(gookitevent.GamerBaseExit, gookitevent.GenGameEvent(gookitevent.SetGamerId(gamerId)))
}

func (p *MsgHandler) OnProcessMsg(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	//ctx := context.Background()
	//log.Info(ctx, "----- OnProcessMsg GetUser:%d", msgque.GetUser())
	//登入的第一包不做阻斷與判斷
	//if msg.Head.Cmd == uint8(proto.SystemCategory_LoginSubsystem) ||
	//	msg.Head.Cmd == uint8(proto.SystemCategory_GSLoginSubsystem) {
	//	return true
	//}

	if msgque.GetUser() != nil {
		//log.Warn(ctx,"----- OnProcessMsg GetUser : %v", msgque.GetUser())
		return true
	}
	//log.Error(ctx, "----- OnProcessMsg GetUser is nil, msgque id:%d", msgque.Id())
	return false
}

func (p *MsgHandler) OnConnectComplete(msgque antnet.IMsgQue, ok bool) bool {
	//ctx := context.Background()
	//log.Info(ctx, "----- OnConnectComplete GetUser:%d", msgque.GetUser())
	return true
}

//func (p *MsgHandler) GetHandlerFunc(msgque antnet.IMsgQue, msg *antnet.Message) antnet.HandlerFunc {
//	if msg.Head == nil {
//		if p.typeMap != nil {
//			if f, ok := p.typeMap[reflect.TypeOf(msg.C2S())]; ok {
//				return f
//			}
//		}
//	} else if r.msgMap != nil {
//		if f, ok := r.msgMap[msg.CmdAct()]; ok {
//			return f
//		}
//	}
//
//	return nil
//}
