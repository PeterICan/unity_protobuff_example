package igamer

import (
	"proto_buffer_example/server/internal/player/data"
	"proto_buffer_example/server/third-party/antnet"
)

type IGamer interface {
	// GetGamerId 取得玩家ID
	GetGamerId() uint64
	// GetMsgque 取得玩家連線模組
	GetMsgque() antnet.IMsgQue
	// EnterGameWorld 玩家物件正式進入遊戲，需要觸發的功能
	EnterGameWorld()
	// GetGamerData 取得玩家資料結構結合
	GetGamerData() *data.GamerData
	// SetMsgQue 設定玩家連線模組
	SetMsgQue(msgque antnet.IMsgQue)
	//// GetBaseData 取得玩家基本資料
	//GetBaseData() data.IGameBase
}
