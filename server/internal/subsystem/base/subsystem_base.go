package base

import (
	"context"
	"proto_buffer_example/server/internal/player/interface/igamer"
	"proto_buffer_example/server/third-party/antnet"
	"proto_buffer_example/server/tools/customize"
)

// Information 每個系統需要包含的資料結構
type Information struct {
	//System        proto.SystemCategory_SystemDefine
	Category uint16
	//該子系統所屬的伺服器ID
	ServerId int32
	//伺服器資訊
	//ServerInfo *info.ServerInfo
}

// ISubsystemBase 所有subsystem需要實現的介面
type ISubsystemBase interface {
	// Init 子系統初始化
	Init(serviceId int32)
	// RegisterMsgHandler 子系統註冊監聽封包
	RegisterMsgHandler(handler *customize.MsgHandler)
	// RegisterMsgParser 子系統註冊封包解析方式
	RegisterMsgParser(parser *antnet.Parser, customParser *customize.JsonRouteParser)
	// Start 子系統啟動
	Start()
	// Stop 子系統停止
	Stop()
	// OnPlayerCreate 當玩家創立帳號
	OnPlayerCreate(gamer igamer.IGamer)
	// OnPlayerEnter 當玩家登入遊戲
	OnPlayerEnter(ctx context.Context, gamer igamer.IGamer)
	// OnPlayerExit 當玩家離開遊戲
	OnPlayerExit(ctx context.Context, gamer igamer.IGamer)
	// OnPlayerLoadFromDB 當玩家資料從DB讀出
	OnPlayerLoadFromDB(ctx context.Context, gamer igamer.IGamer)
	//ListenerBroadcast 監聽廣播事件
	ListenerBroadcast()
}
