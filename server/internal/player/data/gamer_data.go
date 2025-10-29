package data

/*
	☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆
		玩家物件的資料結構集名體

		每次在GamerData新增一種資料結構時
		需要在 CreateGamer,TestCreateGamer 跟 CreateGamerData 這幾處補上初始化的處理
	☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆
*/

const redisCacheTime = 172800 //600秒是為了測試用 //172800 玩家的快取資料存活時間為172800秒 (2天)

// GamerData 玩家資料集合體
type GamerData struct {
	//玩家基礎資料
	*GameBase
	//帳號資料
	//Account *Account
	//位置
	Position *Position
}

// CreateGamerData 產生並初始化一個玩家資料層物件
func CreateGamerData(gamerId uint64) *GamerData {

	//accountId := gamerId

	return &GamerData{
		GameBase: &GameBase{GamerId: gamerId}, //LastLoginTime:  time_helper.TaipeiTime(),
		//LastLogoutTime: time_helper.TaipeiTime(),

		//Account: &Account{AccountId: accountId},
	}
}
