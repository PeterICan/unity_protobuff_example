package player

import (
	"proto_buffer_example/server/internal/player/data"
	"proto_buffer_example/server/third-party/antnet"
	"time"
)

type Gamer struct {
	//玩家資料集合
	*data.GamerData
	//玩家連線模組
	MsgQue antnet.IMsgQue `redis:"-"`
	//封包快取結構 (不進資料庫)
	//*data.MessageCatch `redis:"-"`
	//chenEvent chan *ChannelEvent
	//玩家定時檢查資料異動的計時器
	checkDirtyTimer *time.Timer `redis:"-"`
	//心跳計時器
	heartBeatTimer *time.Timer `redis:"-"`
	//心跳次數
	index int32
}

func (g Gamer) GetMsgque() antnet.IMsgQue {
	//TODO implement me
	panic("implement me")
}

func (g Gamer) EnterGameWorld() {
	//TODO implement me
	panic("implement me")
}

func (g Gamer) GetGamerData() *data.GamerData {
	//TODO implement me
	panic("implement me")
}

func (g Gamer) SetMsgQue(msgque antnet.IMsgQue) {
	//TODO implement me
	panic("implement me")
}
