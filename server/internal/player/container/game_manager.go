package container

// GamerManager 玩家物件管理者
type GamerManager struct {
	GamerContainer
}

// SubsystemManager 不處理玩家物件的子系統管理者(功能伺服器專用，例：聊天、公會、排行榜...)
type SubsystemManager struct {
	SubsystemContainer
}
