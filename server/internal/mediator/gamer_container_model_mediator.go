package mediator

import (
	"context"

	"proto_buffer_example/server/internal/player/interface/igamer"
	"proto_buffer_example/server/third-party/antnet"
)

type IGamerContainerModelMediator interface {
	// GetInstance 取得實體
	GetInstance() IGamerContainerModelMediator
	// GetOnlineGamer 取得線上玩家物件
	GetOnlineGamer(gamerId uint64) (igamer.IGamer, bool)
	// GetMsgQueFromGamerId 透過玩家ID取得msgque
	GetMsgQueFromGamerId(gamerId uint64) (antnet.IMsgQue, bool)
	// GetGamerFromMsgque 傳入msgque取得玩家物件
	GetGamerFromMsgque(msgque antnet.IMsgQue) (igamer.IGamer, bool)
	// GetIGamer 取得玩家物件
	//GetIGamer(gamerId uint64) (igamer.IGamer, define.GamerDataResourceType)
	// IsExists 玩家是否存在
	IsExists(gamerId uint64) bool
	// AttachGamerContainer 將玩家加入容器內暫存管理
	AttachGamerContainer(ctx context.Context, gamer igamer.IGamer, repeatLogin, firstTimeLogin bool, firstLoginToday bool)
	// RemoveContainerGamer 從容器中移除玩家
	RemoveContainerGamer(gamerId uint64)
	// LoadGamerBaseData 讀取玩家的基礎資料並返回玩家物件
	//LoadGamerBaseData(accountData idata.IAccount) (bool, igamer.IGamer)
	// CallGamerFunc 玩家回呼的函式
	CallGamerFunc(f func(gamer igamer.IGamer))
	// NewGamerInitData 新玩家登入時創建基礎資料與取得玩家物件
	//NewGamerInitData(accountData idata.IAccount, loginInfo *proto.Login_ClientLoginInfo) igamer.IGamer
	// GetGamerCount 取得玩家數量
	GetGamerCount() int
	//SaveGamer 儲存玩家資料
	SaveGamer(gamerId uint64, value interface{})
	//SaveIGamer 儲存玩家資料
	SaveIGamer(gamer igamer.IGamer, value ...interface{})
	//KickGamer 踢除玩家
	KickGamer(gamerId uint64)
	//KickAllGamer 踢除所有玩家
	KickAllGamer()
}
