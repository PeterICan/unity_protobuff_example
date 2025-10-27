package data

// IData 玩家各系統資料的儲存介紹
type IData interface {
	// Initialize 初始化資料
	Initialize()
	//// LoadDataFromDB 從DB把資料讀回
	//LoadDataFromDB() error
	//// LoadDataFromCache 從快取把資料讀回 todo 這個目前沒在使用，或許可以刪除
	//LoadDataFromCache() error
	//// LoadDataFromCachePipeline 從快取把資料讀回並使用pipeline機制
	//LoadDataFromCachePipeline(pipeline redis.Pipeliner) error
	// Save 將資料存入DB與快取
	Save()
	// SaveToCache 將資料存入快取
	SaveToCache()
	// SetDirty 資料有變動(會即時寫入快取，等定時器觸發時再寫回DB)
	SetDirty()
	// IsDirty 資料是否有異動
	IsDirty() bool
	// OfflineSave 玩家離線時的資料寫入資料庫與快取流程 (主要是可以在離線儲存時同時通知該系統的計時器應該停止計時)
	OfflineSave()

	/* LoadDataFromDBPipeline 從DB把資料讀回 todo 此函式在針對不同資料庫時無法同個介面處理，先註解
	/LoadDataFromDBPipeline(tx *gorm.DB) error */
}
