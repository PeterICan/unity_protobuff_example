package player

import (
	"fmt"
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

func (p *Gamer) GetMsgque() antnet.IMsgQue {
	return p.MsgQue
}

func (p *Gamer) EnterGameWorld() {
	fmt.Println("EnterGameWorld gamer:", p.GamerId)
	return
	//ctx := context.Background()
	//log.Info(ctx, "EnterGameWorld gamer:%d", p.GamerId)
	////修改玩家登入狀態
	//p.GetBaseData().SetOnline(true)
	////啟動dirty資料檢查
	//p.checkDirtyCallback()
	////啟動玩家心跳
	//p.startHeartbeat()
}

func (p *Gamer) GetGamerData() *data.GamerData {
	return p.GamerData
}

func (p *Gamer) SetMsgQue(msgque antnet.IMsgQue) {
	p.MsgQue = msgque
	p.MsgQue.SetUser(p.GamerId)
}
