package position

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/internal/mediator"
	"proto_buffer_example/server/internal/subsystem/base"
	"proto_buffer_example/server/third-party/antnet"
	"proto_buffer_example/server/tools/customize"
)

/*
	玩家信息子系統
	用於管理玩家的基本信息
	心跳包、升級、獲取玩家信息等
*/

// PositionController 玩家信息子系統
type PositionController struct {
	base.Information
}

// Init 子系統初始化
func (p *PositionController) Init(serviceId int32) {
	//p.Category = uint16(proto.SystemCategory_GamerInfoSubsystem)
	p.ServerId = serviceId
}

// RegisterMsgHandler 子系統註冊監聽封包
func (p *PositionController) RegisterMsgHandler(handler *customize.MsgHandler) {
	handler.RegisterMsg(&json_api.C2SPositionUpdate{}, p.OnC2SPositionUpdate)
}

// RegisterMsgParser 子系統註冊封包解析方式
func (p *PositionController) RegisterMsgParser(parser *antnet.Parser, customParser *customize.JsonRouteParser) {
	customParser.RegisterMsg("position/update", &json_api.C2SPositionUpdate{}, nil)
}

// Start 子系統啟動
func (p *PositionController) Start() {
}

// Stop 子系統停止
func (p *PositionController) Stop() {
}

func (p *PositionController) OnC2SPositionUpdate(msgque antnet.IMsgQue, msg *antnet.Message) bool {

	ctx := context.Background()
	//ctx = log.MsgTraceIntoCtx(ctx, p.ServerId)
	c2s := msg.C2S().(*json_api.C2SPositionUpdate)
	if c2s == nil {
		log.Fatal(ctx, "OnC2SRetrieve c2s is nil, msgqueId:%v getUser:%v", msgque.Id(), msgque.GetUser())
		return false
	}

	gamer, ok := mediator.IGamerContainerModelMdr.GetGamerFromMsgque(msgque)
	if !ok {
		log.Fatal(ctx, "OnC2SRetrieve gamer is nil, msgqueId:%v getUser:%v", msgque.Id(), msgque.GetUser())
		return false
	}

	log.Default().Println("OnC2SPositionUpdate called. gamerId:", gamer.GetGamerId(), "c2s:", c2s)

	rsp := antnet.NewByteHeadlessMsg()
	s2c := &json_api.S2CPositionUpdate{}

	mediator.IPositionModelMdr.OnUpdatePosition(ctx, gamer, c2s, s2c)
	s2cJsonData, err := json.Marshal(s2c)
	if err != nil {
		fmt.Printf("Error marshalling response to JSON: %v\n", err)
		return false
	}
	rsp.Data = s2cJsonData
	msgque.Send(rsp)
	//gamer.AddCatchMessage(rsp)
	//ctx = log.MsgUpdateTraceTime(ctx)
	//log.Trace(ctx, "OnC2SRetrieve gamer:%d s2c:%v", gamer.GetGamerId(), s2c)
	return true
}
