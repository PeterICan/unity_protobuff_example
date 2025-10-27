package data

import (
	"github.com/go-redis/redis"
	time2 "proto_buffer_example/server/tools/time_helper"
	"time"
)

const (
	NumberRangeEnergy = 100000 //能量最大上限
)

// GameBase 玩家基礎資料
type GameBase struct {
	//玩家ID
	GamerId uint64 `gorm:"column:id;primaryKey" bson:"_id"`
	//玩家等級
	Level int32 `gorm:"default:1;not null" bson:"Level"`
	//玩家登入的遊戲伺服器編號
	GameServerId int32 `gorm:"default:0;not null" bson:"GameServerId"`
	//最後登入時間
	LastLoginTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" bson:"LastLoginTime"`
	//倒數第二次登入時間
	SecondsLastLoginTime int64 `gorm:"not null;default:CURRENT_TIMESTAMP" bson:"SecondsLastLoginTime"`
	//最後登出時間
	LastLogoutTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" bson:"LastLogoutTime"`
	//玩家總遊玩時間(秒)
	TotalInGameTime int64 `gorm:"default:0;not null" bson:"TotalInGameTime"`
	//玩家是否在線
	isOnline bool `gorm:"-" bson:"-"`
	//測試資料組別
	TestGroup int32 `bson:"TestGroup"`
	//是否通過開頭
	IsOpeningPassed bool `bson:"IsOpeningPassed"`
	//第一次購買時間
	FirstPurchaseTime int64 `bson:"FirstPurchaseTime"`
	//資料dirty狀態 此欄位不進資料庫
	dirty bool `gorm:"-" redis:"-" bson:"-"`
}

func (p *GameBase) Initialize() {
	p.Level = 1
	p.LastLoginTime = time2.TaipeiTime()
	p.LastLogoutTime = time.Time{}
	p.TotalInGameTime = 0
	//p.Energy = p.getInitialEnergyDefault()
	p.isOnline = true
	//p.TestGroup = int32(p.GamerId % uint64(modulusdata.GetInstance().GetModulusDataFromType(modulusdata.ABTestGroupCount).Value))
	//p.InitEnergyRecoveryFunc(p.getInitialMaxEnergy(), p.getRecoveryEnergySeconds(), p.getRecoveryEnergies(), p.getEnergyLimit(), p.SetDirty)
	p.SetDirty()
}

func (p *GameBase) GetGamerId() uint64 {
	return p.GamerId
}

// GetTestGroup 取得測試群組
func (p *GameBase) GetTestGroup() int32 {
	return p.TestGroup
}

func (p *GameBase) GetLevel() int32 {
	return p.Level
}

func (p *GameBase) IsOnline() bool {
	return p.isOnline
}

func (p *GameBase) SetOnline(online bool) {
	p.isOnline = online
	//if !p.isOnline { //離線時，停止計時器
	//	p.StopTimer()
	//}
}

// UpdateLastLoginTime 更新玩家最後上線時間
func (p *GameBase) UpdateLastLoginTime() {
	p.SecondsLastLoginTime = p.LastLoginTime.Unix()
	p.LastLoginTime = time2.TaipeiTime()
	p.SetDirty()
}

// GetLastLoginTime 取得玩家最後上線時間
func (p *GameBase) GetLastLoginTime() time.Time {
	return p.LastLoginTime
}

// UpdateLastLogoutTime 刷新登出時間，通知更新總遊戲時間
func (p *GameBase) UpdateLastLogoutTime() int64 {
	p.LastLogoutTime = time2.TaipeiTime()
	duration := p.LastLogoutTime.Sub(p.LastLoginTime).Seconds()
	p.TotalInGameTime = p.TotalInGameTime + int64(duration)
	p.SetDirty()
	return int64(duration)
}

// GetTotalInGameTime 線上玩家取得即時的總遊戲時間()
func (p *GameBase) GetTotalInGameTime() int64 {
	duration := time2.TaipeiTime().Sub(p.LastLoginTime).Seconds()
	return p.TotalInGameTime + int64(duration)
}

func (p *GameBase) GetSecondLastLoginTime() int64 {
	return p.SecondsLastLoginTime
}

//func (p *GameBase) LoadDataFromDBPipeline(tx *gorm.DB) error {
//	context.Background()
//	//ctx := context.Background()
//	//err := tx.Take(p, "id = ?", p.GamerId).Error
//	//if err != nil && err != gorm.ErrRecordNotFound {
//	//	log.Error(ctx, "LoadDataFromDBPipeline p:%v err:%v", p, err)
//	//	tx.Rollback()
//	//	return err
//	//}
//	return nil
//}

func (p *GameBase) LoadDataFromCachePipeline(pipeline redis.Pipeliner) error {

	panic("not implemented")
	//ctx := context.Background()
	//err := GetGamerHashPipeline(p.GamerId, p, pipeline)
	//if err != nil {
	//	log.Error(ctx, "LoadDataFromCachePipeline p:%v err:%v", p, err)
	//}
	//return err
}

func (p *GameBase) OfflineSave() {
	p.Save()
}

func (p *GameBase) Save() {
	//ctx := context.Background()
	//now := time.Now()
	panic("not implemented")
	p.SaveToCache()
	//err := maindb.DBHelper.Save(p.GamerId, p)
	//if err != nil {
	//	log.Error(ctx, "gamer:%d Save err:%v", p.GamerId, err)
	//	return
	//}

	p.dirty = false
	//since_counter.SinceCounter(ctx, now)
}

func (p *GameBase) SaveToCache() {
	panic("not implemented")
	//ctx := context.Background()
	//err := SetGamerHash(p.GamerId, p)
	//if err != nil {
	//	log.Error(ctx, "gamer:%d SaveToCache err:%v", p.GamerId, err)
	//}
}

func (p *GameBase) SetDirty() {
	//log.Info(ctx, "---------- SetDirty gamer:%d", p.GamerId)
	p.dirty = true
	p.SaveToCache()
}

func (p *GameBase) IsDirty() bool {
	return p.dirty
}
