package position

import (
	"context"
	"encoding/json"
	"fmt"
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/internal/mediator"
	"proto_buffer_example/server/internal/player/interface/igamer"
	"proto_buffer_example/server/internal/subsystem/base"
	"proto_buffer_example/server/third-party/antnet"
)

func NewPositionModel() *PositionModel {
	p := &PositionModel{}
	mediator.IPositionModelMdr = p
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
	gamer.GetPositionData().Initialize()
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

// OnUpdatePosition 玩家位置更新
func (p *PositionModel) OnUpdatePosition(ctx context.Context, gamer igamer.IGamer, c2s *json_api.C2SPositionUpdate, s2c *json_api.S2CPositionUpdate) {
	gamer.GetPositionData().SetPosition(c2s.GetX(), c2s.GetY(), c2s.GetZ())
	s2c.Route = "position/update"
	s2c.Error = nil

	notify := antnet.NewByteHeadlessMsg()
	notifyS2C := &json_api.S2CNotifyWorldPositionChange{}
	notifyS2C.Route = "position/notify_world_position_change"
	notifyS2C.Positions = p.getOnlinePlayerPosition()
	s2cJson, err := json.Marshal(notifyS2C)
	if err != nil {
		fmt.Printf("Error marshalling response to JSON: %v\n", err)
		return
	}
	notify.Data = s2cJson
	antnet.SendGroup("BaseGroup", notify)
}

func (p *PositionModel) getOnlinePlayerPosition() []*json_api.WorldPosition {
	onlinePlayers := mediator.IGamerContainerModelMdr.GetAllOnlineGamers()
	positions := make([]*json_api.WorldPosition, len(onlinePlayers))
	for index, gamer := range onlinePlayers {
		PlayerId :=
			fmt.Sprintf("%d", gamer.GetGamerId())
		posX, posY, posZ := gamer.GetPositionData().GetPosition()
		position := &json_api.WorldPosition{
			PlayerId: PlayerId,
			X:        posX,
			Y:        posY,
			Z:        posZ,
		}
		positions[index] = position
	}

	return positions
}

func (p *PositionModel) groupNotifyDailyReset() {
	//s2c := &proto.GamerInfo_S2CDailyResetNotify{Result: proto.GamerInfo_S2CDailyResetNotify_Success}
	//notify := antnet.NewTagMsg(uint8(p.Category), uint8(proto.GamerInfoAct_DailyResetNotify), 0)
	//notify.Data = antnet.PbData(s2c)
	//antnet.SendGroup(define.BaseGroup, notify)
	//log.Info(context.Background(), "groupNotifyDailyReset group:Scene s2c:%v", s2c)
}
