package base

import (
	"context"
	"proto_buffer_example/server/internal/player"
	"proto_buffer_example/server/internal/player/interface/igamer"
	"proto_buffer_example/server/third-party/antnet"
	"proto_buffer_example/server/tools/customize"
)

type IPlayerManagerGet interface {
	// GetOnlineGamer 取得傳入的id玩家物件
	GetOnlineGamer(gamerId uint64) (*player.Gamer, bool)
	// GetRandGamer 取得隨機的線上玩家物件，但不包含傳入的玩家id
	GetRandGamer(gamerId uint64) (*player.Gamer, bool)
}

// ISubsystemModel 所有子系統Model必需實作介面
type ISubsystemModel interface {
	// Init manager 初始化
	Init(serviceId int32)
	// Start manager 第一次啟動
	Start()
	// Stop manager 停止運行
	Stop()
	// SaveToDB 儲存系統資料進DB
	SaveToDB()
	// LoadFromDB 將DB資料讀回至系統
	LoadFromDB()
	// OnPlayerCreate 當玩家創立帳號
	OnPlayerCreate(ctx context.Context, gamer igamer.IGamer)
	// OnPlayerEnter 當玩家登入遊戲
	OnPlayerEnter(ctx context.Context, gamer igamer.IGamer)
	// OnPlayerExit 當玩家離開遊戲
	OnPlayerExit(ctx context.Context, gamer igamer.IGamer)
	// OnPlayerLoadFromDB 當玩家資料從DB讀出
	OnPlayerLoadFromDB(ctx context.Context, gamer igamer.IGamer)
}

// ISubsystemController 所有子系統Controller必需實現介面
type ISubsystemController interface {
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
}
