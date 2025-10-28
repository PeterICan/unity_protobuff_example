package position

import (
	"context"
	"proto_buffer_example/server/internal/player/interface/igamer"
	"proto_buffer_example/server/internal/subsystem/base"
)

func NewPositionModel() *PositionModel {
	p := &PositionModel{}
	//mediator.IPositionModelMdr = p //TODO
	return p
}

// PositionModel 樣版 model
type PositionModel struct {
	base.Information
}

// Init Model 初始化
func (p *PositionModel) Init(serviceId int32) {
	p.ServerId = serviceId
}

// Start Model 啟動
func (p *PositionModel) Start() {

}

// Stop Model 停止運行
func (p *PositionModel) Stop() {
}

// SaveToDB 儲存系統資料進DB
func (p *PositionModel) SaveToDB() {
}

// LoadFromDB 將DB資料讀回至系統
func (p *PositionModel) LoadFromDB() {
}

// OnPlayerCreate 當玩家創立帳號
func (p *PositionModel) OnPlayerCreate(ctx context.Context, gamer igamer.IGamer) {
}

// OnPlayerEnter 當玩家登入遊戲
func (p *PositionModel) OnPlayerEnter(ctx context.Context, gamer igamer.IGamer) {

}

// OnPlayerExit 當玩家離開遊戲
func (p *PositionModel) OnPlayerExit(ctx context.Context, gamer igamer.IGamer) {

}

// OnPlayerLoadFromDB 當玩家資料從DB讀出
func (p *PositionModel) OnPlayerLoadFromDB(ctx context.Context, gamer igamer.IGamer) {

}

//// OnRetrieve 獲取玩家資料
//func (p *PositionModel) OnRetrieve(ctx context.Context, gamer igamer.IGamer, s2c *proto.GamerInfo_S2CRetrieve) {
//	panic("implement me")
//}

func (p *PositionModel) groupNotifyDailyReset() {
	//s2c := &proto.GamerInfo_S2CDailyResetNotify{Result: proto.GamerInfo_S2CDailyResetNotify_Success}
	//notify := antnet.NewTagMsg(uint8(p.Category), uint8(proto.GamerInfoAct_DailyResetNotify), 0)
	//notify.Data = antnet.PbData(s2c)
	//antnet.SendGroup(define.BaseGroup, notify)
	//log.Info(context.Background(), "groupNotifyDailyReset group:Scene s2c:%v", s2c)
}
