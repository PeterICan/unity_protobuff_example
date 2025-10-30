package data

import (
	"fmt"
	"proto_buffer_example/server/internal/persistence"
)

type Position struct {
	GamerId uint64

	PosX float32
	PosY float32
	PosZ float32

	dirty bool
}

////////////IData//////////////////////////////

func (p *Position) Initialize() {
	fmt.Println("Position Initialize called gamerId:", p.GamerId)
	p.PosX = 0
	p.PosY = 0
	p.PosZ = 0
	p.SetDirty()
}

func (p *Position) Save() {
	fmt.Println("Position Save called gamerId:", p.GamerId)
	p.SaveToCache()
	persistence.SavePlayerData(p.GamerId, p)
	//err := maindb.DBHelper.Save(p.GamerId, p)
	p.dirty = false
	//since_counter.SinceCounter(ctx, now)
}

func (p *Position) SaveToCache() {
	fmt.Println("Position SaveToCache called: ", p.GamerId)
}

func (p *Position) SetDirty() {
	p.dirty = true
	p.SaveToCache()
	p.Save()
}

func (p *Position) IsDirty() bool {
	return p.dirty
}

func (p *Position) OfflineSave() {
	fmt.Println("Position OfflineSave called: ", p.GamerId)
	p.Save()
}

////////////IPosition//////////////////////////////

func (p *Position) SetPosition(x, y, z float32) {
	p.PosX = x
	p.PosY = y
	p.PosZ = z
	p.SetDirty()
}

func (p *Position) GetPosition() (float32, float32, float32) {
	return p.PosX, p.PosY, p.PosZ
}
