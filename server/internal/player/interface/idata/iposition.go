package idata

import "proto_buffer_example/server/internal/player/data"

// IPosition 玩家位置資料介面
type IPosition interface {
	data.IData

	// SetPosition 設定玩家位置
	SetPosition(x, y, z float32)
	// GetPosition 取得玩家位置
	GetPosition() (float32, float32, float32)
}
