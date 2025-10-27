package player

//// CreateGamer 傳入Account資料並創建一個玩家物件
//func CreateGamer(accountData idata.IAccount, serverId int32) *Gamer {
//	//gamerId 跟 accountId 是同一個數值
//	gamerId := accountData.GetAccountId()
//	if gamerId == 0 {
//		fmt.Printf("CreateGamer gamerId:%d serverId:%d error", gamerId, serverId)
//		return nil
//	}
//	//log.Info(ctx, "CreateGamer gamer:%d serverId:%d", gamerId, serverId)
//	gamer := &Gamer{
//		GamerData: &data.GamerData{
//			GameBase: &data.GameBase{GamerId: gamerId,
//				LastLoginTime:  time_helper.TaipeiTime(),
//				LastLogoutTime: time_helper.TaipeiTime(),
//			},
//			Account: accountData.(*data.Account),
//		},
//		//MessageCatch: data.NewMessageCatchData(gamerId),
//	}
//	return gamer
//}
