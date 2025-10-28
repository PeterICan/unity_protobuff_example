package container

import (
	"context"
	"log"
	"proto_buffer_example/server/internal/mediator"
	"proto_buffer_example/server/internal/persistence"
	"proto_buffer_example/server/internal/player"
	"proto_buffer_example/server/internal/player/interface/igamer"
	"proto_buffer_example/server/third-party/antnet"
	"proto_buffer_example/server/tools"
	"sync"
	"sync/atomic"
)

// GamerContainer 玩家物件容器與管理模組
type GamerContainer struct {
	SubsystemContainer
	MsgqueMap map[uint64]antnet.IMsgQue
	GamerMap  map[uint64]*player.Gamer

	mapLock     sync.RWMutex
	saveChannel chan interface{}

	saveFlag int32 //新增一個 saveFlag 變數，用於標記保存標誌
	//所屬的伺服器id
	serverId int32
}

func (p *GamerContainer) InitGamerContainer(serverId int32) {
	mediator.IGamerContainerModelMdr = p
	p.saveChannel = make(chan interface{})
	p.serverId = serverId
	p.MsgqueMap = make(map[uint64]antnet.IMsgQue)
	p.GamerMap = make(map[uint64]*player.Gamer)

}

// InitRegisterSub 初始化 Redis 訂閱監控相關(需要最後初始化)
func (p *GamerContainer) InitRegisterSub() {
	//ctx := context.Background()

	//gookitevent.GetInstance().On(gookitevent.SubsystemManager.String(), gookitevent.ScheduleCycleStart, event.ListenerFunc(p.dailyRefreshHandler), event.Normal)
	//gookitevent.GetInstance().On(gookitevent.SubsystemManager.String(), gookitevent.ScheduleCycleStop, event.ListenerFunc(p.dailyResetLogHandler), event.Normal)
}

func (p *GamerContainer) GetInstance() mediator.IGamerContainerModelMediator {
	return p
}

func (p *GamerContainer) AttachGamerContainer(ctx context.Context, gamer igamer.IGamer, repeatLogin, firstTimeLogin bool, firstLoginToday bool) {
	//if !repeatLogin {
	p.mapLock.Lock()
	p.MsgqueMap[gamer.GetGamerId()] = gamer.GetMsgque()
	p.GamerMap[gamer.GetGamerId()] = gamer.(*player.Gamer)
	p.mapLock.Unlock()
	//}

	// 註冊玩家MQ群組
	p.registerMsgqueGroup(ctx, gamer)

	//通知各子系統去載入該玩家的資料
	//p.onGamerLoadFromDB(ctx, gamer)

	//如果玩家是新創帳號時，觸發所有子系統的創建玩家事件
	if firstTimeLogin {
		p.OnGamerCreate(ctx, gamer)
	}
	//玩家登入紀錄(用於統計留存率)
	//p.loginRecord(ctx, gamer, p.serverId)

	//玩家物件的登入世界觸發事件(啟動心跳檢查)
	gamer.EnterGameWorld()
	//觸發所有子系統的登入遊戲世界事件
	p.OnGamerEnter(ctx, gamer)
}

// loginRecord 登入紀錄
//func (p *GamerContainer) loginRecord(ctx context.Context, gamer igamer.IGamer, serverId int32) {
//	//log.Trace(ctx, "loginRecord gamer:%v", gamer.GetGamerId())
//	//logRecord := info.GenerateLoginRecord(gamer.GetAccountData().GetAccountId(), time.Unix(gamer.GetAccountData().GetRegisterTime(), 0))
//	//err := mongodbHelper.GetDBHelper().SaveLog(logRecord.ID, logRecord)
//	//logdefine.CheckLogError(ctx, err)
//
//	//p.gsLoginRecord(ctx, gamer, serverId)
//
//	//gamermeasurment.WPGamerLogin(gamer.GetAccountData().GetRegisterCountry(), gamer.GetBaseData().GetLevel())
//
//	//redishelper.LoginRecordRetentionRate(gamer.GetAccountData().GetAccountId(), time.Unix(gamer.GetAccountData().GetRegisterTime(), 0))
//}

// gsLoginRecord 登入紀錄(使用mysql來儲存game log)
//func (p *GamerContainer) gsLoginRecord(ctx context.Context, gamer igamer.IGamer, serverId int32) {
//
//	firstPurchaseTime := gamer.GetBaseData().GetFirstPurchaseTime()
//	if firstPurchaseTime > 0 {
//		//gsLoginLog.FirstPurchaseTime = time_helper.ParseUnixTimeToTaipeiTomeZonePtr(firstPurchaseTime)
//	}
//	err := dbhelper.GetLogDBHelper().SaveLog(gsLoginLog)
//	logdefine.CheckLogError(ctx, err)
//}

func (p *GamerContainer) registerMsgqueGroup(ctx context.Context, gamer igamer.IGamer) {
	// 所有玩家都要新增的群組
	//define.BaseGroup
	gamer.GetMsgque().SetGroupId("BaseGroup")

	//log.Info(ctx, "registerMsgqueGroup:%v gamer:%d", "Procure", gamer.GetGamerId())
}

func (p *GamerContainer) RemoveContainerGamer(gamerId uint64) {
	ctx := context.Background()
	gamer, ok := p.getGamer(gamerId, true)
	if !ok {
		//log.Error(ctx, "RemoveContainerGamer GamerSyncMap can't find gamer:%d", gamerId)
	} else {
		p.GamerQuitProcess(ctx, gamer)
	}
	log.Default().Println("移除玩家連線 編號:", gamerId)
	//log.Info(ctx, "RemoveContainerGamer gamer:%d", gamerId)

	p.mapLock.Lock()
	gamer, ok = p.GamerMap[gamerId]
	if ok {
		if p.GamerMap[gamerId].GetMsgque() != nil {
			delete(p.GamerMap, gamerId)
		}
		if p.MsgqueMap[gamerId] != nil {
			delete(p.MsgqueMap, gamerId)
		}
		//if !gamer.GetBaseData().IsOnline() {
		//delete(p.MsgqueMap, gamerId)
		//delete(p.GamerMap, gamerId)
		//}
	}
	p.mapLock.Unlock()
}

func (p *GamerContainer) GamerQuitProcess(ctx context.Context, gamer igamer.IGamer) {
	p.onGamerExit(ctx, gamer)
	//找的到玩家物件再執行以下動作
	//stayTime := gamer.QuitGameWorld()
	//登出紀錄
	//p.logoutLog(gamer, p.serverId, stayTime)
	//p.logoutRecord(context.Background(), gamer, p.serverId, stayTime)
	//gamermeasurment.WPGamerLogout(gamer.GetAccountData().GetRegisterCountry(), gamer.GetBaseData().GetLevel(), stayTime)
}

//func (p *GamerContainer) logoutLog(gamer igamer.IGamer, serverId int32, stayTime int64) {
//	tags := gamer.GetPointTags(serverId)
//	fields := gamer.GetPointFieldsAny(influxdefine.StayTimeField.String(), stayTime)
//	fields[influxdefine.ValueField.String()] = 1
//	fields[influxdefine.CCUField.String()] = p.GetGamerCount()
//	point := influxhelper.GenPoint(influxdefine.GamerLogout.String(), tags, fields)
//	influxhelper.GetInfluxHelper().WritePoint(point)
//}

//// gsLoginRecord 登入紀錄(使用mysql來儲存game log)
//func (p *GamerContainer) logoutRecord(ctx context.Context, gamer igamer.IGamer, serverId int32, stayTime int64) {
//	logoutLog := &logdefine.LogoutLog{
//		SdkId:           gamer.GetAccountData().GetUUID(),
//		AccountId:       gamer.GetAccountData().GetAccountId(),
//		Level:           gamer.GetBaseData().GetLevel(),
//		ServerID:        serverId,
//		StayTime:        stayTime,
//		RegisterCountry: gamer.GetAccountData().GetRegisterCountry(),
//	}
//	firstPurchaseTime := gamer.GetBaseData().GetFirstPurchaseTime()
//	if firstPurchaseTime > 0 {
//		logoutLog.FirstPurchaseTime = time_helper.ParseUnixTimeToTaipeiTomeZonePtr(firstPurchaseTime)
//	}
//	err := dbhelper.GetLogDBHelper().SaveLog(logoutLog)
//	logdefine.CheckLogError(ctx, err)
//}

func (p *GamerContainer) CallGamerFunc(f func(gamer igamer.IGamer)) {
	p.mapLock.Lock()
	defer p.mapLock.Unlock()

	//ctx := context.Background()
	//now := time.Now()
	for _, gamer := range p.GamerMap {
		//log.Warn(ctx, "CallGamerFunc gamerId::%v", gamer.GetGamerId())
		f(gamer)
	}

	//since_counter.SinceCounter(ctx, now)
}

//// OldLoadGamerBaseData 讀取玩家此伺服器的GameBaseData
//func (p *GamerContainer) OldLoadGamerBaseData(accountData idata.IAccount) (bool, igamer.IGamer) {
//	ctx := context.Background()
//	gamerId := accountData.GetGSGamerId(p.serverId)
//	if gamerId == 0 {
//		log.Panic(ctx, "LoadGamerBaseData GetGSGamerId error accountData:%v serverId:%d", accountData, p.serverId)
//		return false, nil
//	}
//	gamer := player.CreateGamer(accountData, p.serverId)
//
//	err := gamer.GetBaseData().LoadDataFromDB()
//	if err != nil {
//		log.Panic(ctx, "LoadGamerBaseData GetBaseData load error:%v serverId:%d", err, p.serverId)
//	}
//
//	err = gamer.GetDeviceData().LoadDataFromDB()
//	if err != nil {
//		log.Panic(ctx, "LoadGamerBaseData GetDeviceData load error:%v serverId:%d", err, p.serverId)
//	}
//	return true, gamer
//}

//// LoadGamerBaseData 非阻塞式讀取玩家此伺服器的GameBaseData
//func (p *GamerContainer) LoadGamerBaseData(accountData idata.IAccount) (bool, igamer.IGamer) {
//	ctx := context.Background()
//	gamerId := accountData.GetGSGamerId(p.serverId)
//	if gamerId == 0 {
//		log.Panic(ctx, "LoadGamerBaseData GetGSGamerId error accountData:%v serverId:%d", accountData, p.serverId)
//		return false, nil
//	}
//	log.Info(ctx, "LoadGamerBaseData gamerId:%v", gamerId)
//	gamer := player.CreateGamer(accountData, p.serverId)
//	// 使用 goroutine 並發執行讀取資料的操作
//	errChan := make(chan error, 2)
//	go func() {
//		errChan <- gamer.GetBaseData().LoadDataFromDB()
//	}()
//	go func() {
//		errChan <- gamer.GetDeviceData().LoadDataFromDB()
//	}()
//	// 等待兩個 goroutine 完成讀取操作
//	for i := 0; i < 2; i++ {
//		if err := <-errChan; err != nil {
//			log.Panic(ctx, "LoadGamerBaseData load error:%v serverId:%d accountData:%v", err, p.serverId, accountData)
//			return false, nil
//		}
//	}
//	return true, gamer
//}

func (p *GamerContainer) NewGamerInitData(msgque antnet.IMsgQue) igamer.IGamer {
	nextGamerId := persistence.GetNextGamerId()
	ctx := context.Background()
	//新玩家登入
	gamer := player.CreateGamer(nextGamerId, p.serverId)
	if gamer == nil {
		log.Panic(ctx, "NewGamerInitData gamer:%d", gamer.GetGamerId())
	}
	//log.Info(ctx, "NewGamerInitData gamer:%d", gamer.GetGamerId())

	// 這三個參數以login-server賦值為準
	//gamer.GetAccountData().SetUUID(loginInfo.Uuid)
	//gamer.GetAccountData().SetLoginWay(loginInfo.LoginWay)
	//gamer.GetAccountData().SetOpenId(loginInfo.OpenId)
	//gamer.GetAccountData().SetRegisterTime()
	//gamer.GetAccountData().CheckFirstTimeLoginGameServer(ctx, p.serverId)
	//創建玩家帳號資料
	//gamer.GetAccountData().SetDirty()

	//gamer.GetBaseData().Initialize()

	//err := gamer.GetDeviceData().LoadDataFromDB()
	//if err != nil {
	//	log.Error(ctx, "NewGamerInitData GetDeviceData load error:%v serverId:%d", err, p.serverId)
	//}
	gamer.SetMsgQue(msgque)
	p.AttachGamerContainer(ctx, gamer, false, true, true)

	log.Println("新建玩家:", gamer.GetGamerId())
	return gamer
}

func (p *GamerContainer) onGamerLoadFromDB(ctx context.Context, gamer igamer.IGamer) {
	for i := range p.ModelList {
		p.ModelList[i].OnPlayerLoadFromDB(ctx, gamer)
	}
}

func (p *GamerContainer) OnGamerCreate(ctx context.Context, gamer igamer.IGamer) {
	for i := range p.ModelList {
		p.ModelList[i].OnPlayerCreate(ctx, gamer)
	}
}

func (p *GamerContainer) OnGamerEnter(ctx context.Context, gamer igamer.IGamer) {
	//log.Info(ctx, "OnGamerEnter gamer:%d start", gamer.GetGamerId())
	for i := range p.ModelList {
		//log.Info(ctx, "OnGamerEnter model:%v gamer:%d", p.ModelList[i], gamer.GetGamerId())
		p.ModelList[i].OnPlayerEnter(ctx, gamer)
	}
	//log.Info(ctx, "OnGamerEnter gamer:%d done", gamer.GetGamerId())
}

func (p *GamerContainer) onGamerExit(ctx context.Context, gamer igamer.IGamer) {
	for i := range p.ModelList {
		p.ModelList[i].OnPlayerExit(ctx, gamer)
	}
}

func (p *GamerContainer) GetGamerFromMsgque(msgque antnet.IMsgQue) (igamer.IGamer, bool) {
	gamerId, ok := tools.Interface2uint64(msgque.GetUser())
	if !ok {
		return nil, false
	}

	p.mapLock.RLock()
	gamer, ok := p.GamerMap[gamerId]
	p.mapLock.RUnlock()
	if !ok {
		return nil, ok
	}
	return gamer, ok
}

func (p *GamerContainer) GetMsgQueFromGamerId(gamerId uint64) (antnet.IMsgQue, bool) {
	p.mapLock.RLock()
	msgQue, ok := p.MsgqueMap[gamerId]
	p.mapLock.RUnlock()
	if !ok {
		return nil, ok
	}
	return msgQue.(antnet.IMsgQue), ok
}

func (p *GamerContainer) IsExists(gamerId uint64) bool {
	_, ok := p.getGamer(gamerId, true)
	return ok
}

func (p *GamerContainer) GetGamer(gamerId uint64) (igamer.IGamer, bool) {
	return p.getGamer(gamerId, false)
}

// GetIGamer 取得玩家物件(目前只能取得線上玩家)
func (p *GamerContainer) GetIGamer(gamerId uint64) (igamer.IGamer, bool) {

	if gamer, isOnlinePlayer := p.GetGamer(gamerId); isOnlinePlayer {
		return gamer, true
	}

	////再從redis找起
	//gamerData, err := data.LoadGamerDataFromCachePipeline(gamerId)
	//if err == nil {
	//	gamer := &player.Gamer{
	//		GamerData: gamerData,
	//	}
	//	return gamer, define.GamerDataFromCache
	//} else {
	//	log.Error(ctx, "GetIGamer LoadGamerDataFromCachePipeline gamer:%d err:%v", gamerId, err)
	//}
	//
	//if gamerData, err := data.LoadGamerDataFromDB(gamerId); err == nil {
	//	gamer := &player.Gamer{
	//		GamerData: gamerData,
	//	}
	//	if gamer.GetAccountData().GetAccountId() <= 0 {
	//		return nil, define.GamerDataNotFound
	//	}
	//	//把DB的資料再寫入redis快取裡
	//	err = data.SetGamerHash(gamerId, gamerData)
	//	if err != nil {
	//		log.Error(ctx, "GetIGamer gamer:%d SetGamerHash err:%v", gamerId, err)
	//	}
	//	log.Info(ctx, "GetIGamer gamer:%d load Data from DB, set into redis", gamerId)
	//	return gamer, define.GamerDataFromDB
	//}

	return nil, false
}

func (p *GamerContainer) GetOnlineGamer(gamerId uint64) (igamer.IGamer, bool) {
	return p.getGamer(gamerId, true)
}

func (p *GamerContainer) getGamer(gamerId uint64, needOnline bool) (igamer.IGamer, bool) {
	p.mapLock.RLock()
	gamer, ok := p.GamerMap[gamerId]
	p.mapLock.RUnlock()
	if !ok {
		return nil, ok
	}

	//if needOnline && !gamer.GetBaseData().IsOnline() {
	//	return nil, false
	//}
	return gamer, ok
}

// saveData 儲存gamer資料到資料庫與快取中
func (p *GamerContainer) saveData(gamerId uint64, saveData interface{}) {
	panic("implement me")
	//ctx := context.Background()
	//err := maindb.DBHelper.Save(gamerId, saveData)
	//if err != nil {
	//	//log.Error(ctx, "saveData - gamer:%d channel db save err:%v", gamerId, err)
	//	return
	//}
	//err = data.SetGamerHash(gamerId, saveData)
	//if err != nil {
	//	//log.Error(ctx, "saveData - gamer:%d channel redis save err:%v", gamerId, err)
	//}
	atomic.StoreInt32(&p.saveFlag, 0) //保存完畢，將標誌更新為 0
}

// SaveGamer 儲存gamer數據，value為玩家數據
func (p *GamerContainer) SaveGamer(gamerId uint64, value interface{}) {
	//ctx := context.Background()
	p.mapLock.RLock()
	_, ok := p.GamerMap[gamerId]
	p.mapLock.RUnlock()
	//log.Warn(ctx, "SaveGamer - gamer:%d channel save", gamerId)
	//先進行原子操作，避免其他 goroutine 干擾
	if atomic.CompareAndSwapInt32(&p.saveFlag, 0, 1) {
		p.saveChannel <- value
		go func() {
			saveData := <-p.saveChannel //從 channel 中讀取數據
			if saveData != nil {
				gamerId := gamerId
				_, ok := p.GamerMap[gamerId]
				if ok {
					//如果在線上，以玩家記憶體的資料當成儲存依據
					saveData = p.GamerMap[gamerId]
				}
				p.saveData(gamerId, saveData)
			}
		}()
	}
	if ok { //玩家如果在線上的話，要確保一次只有一個人能覆寫標案
		go func() {
			//log.Warn(ctx, "SaveGamer gamer:%d start channel", gamerId)
			p.saveChannel <- value
		}()
	}
	//將標誌更新為 0，以使其他 goroutine 可以繼續進行 SaveGamer 操作
	atomic.StoreInt32(&p.saveFlag, 0)
}

//func (p *GamerContainer) SaveGamer(gamerId uint64, value interface{}) {
//	ctx := context.Background()
//	p.mapLock.RLock()
//	_, ok := p.GamerMap[gamerId]
//	p.mapLock.RUnlock()
//
//	//先進行原子操作，避免其他 goroutine 干擾
//	if atomic.CompareAndSwapInt32(&p.saveFlag, 0, 1) {
//		p.saveChannel <- value
//		go p.saveData()
//	}
//
//	if ok { //玩家如果在線上的話，要確保一次只有一個人能覆寫標案
//		go func() {
//			//log.Warn(ctx, "SaveGamer gamer:%d start channel", gamerId)
//			p.saveChannel <- value
//			//如果在線上，是否要以玩家記憶體的資料當成儲存依據? 而不是傳進來的data ?
//		}()
//	} else {
//		err := maindb.DBHelper.Save(gamerId, value)
//		if err != nil {
//			log.Error(ctx, "SaveGamer offline gamer:%d offline save err:%v", gamerId, err)
//			return
//		}
//		err = data.SetGamerHash(gamerId, value)
//		if err != nil {
//			log.Error(ctx, "SaveGamer offline gamer:%d offline SetGamerHash err:%v", gamerId, err)
//		}
//		return
//	}
//	select {
//	case saveData := <-p.saveChannel:
//		//log.Info(ctx, "SaveGamer %v", saveData)
//		err := maindb.DBHelper.Save(gamerId, saveData)
//		if err != nil {
//			log.Error(ctx, "SaveGamer - gamer:%d channel save err:%v", gamerId, err)
//			return
//		}
//		err = data.SetGamerHash(gamerId, saveData)
//		if err != nil {
//			log.Error(ctx, "SaveGamer - gamer:%d channel SetGamerHash err:%v", gamerId, err)
//		}
//	}
//}

func (p *GamerContainer) SaveIGamer(gamer igamer.IGamer, value ...interface{}) {
	//ctx := context.Background()
	//log.Info(ctx, "SaveIGamer - gamer:%d channel save", gamer.GetGamerId())
	p.SaveGamer(gamer.GetGamerId(), value)
}

func (p *GamerContainer) KickAllGamer() {
	//ctx := context.Background()
	//log.Warn(ctx, "KickAllGamer ... Start")
	p.mapLock.Lock()

	var wg sync.WaitGroup
	wg.Add(len(p.GamerMap))

	for _, gamer := range p.GamerMap {
		go func(gamer igamer.IGamer) {
			defer wg.Done()
			//log.Warn(ctx, "KickAllGamer gamerId::%v", gamer.GetGamerId())
			gamer.GetMsgque().Stop()
		}(gamer)
	}

	p.mapLock.Unlock()
	// Wait for all gamers' data to be saved
	wg.Wait()
	//log.Warn(ctx, "KickAllGamer ... Done")
}

func (p *GamerContainer) KickGamer(gamerId uint64) {
	//ctx := context.Background()
	gamer, ok := p.GetGamer(gamerId)
	if ok {
		//log.Info(ctx, "KickGamer gamerId::%v", gamerId)
		gamer.GetMsgque().Stop()
	} else {
		//log.Warn(ctx, "KickGamer can't find gamerId::%v", gamerId)
	}
}

//func (p *GamerContainer) OnGamerMsgQueExit(msg *redishelper.PublishMessage) bool {
//	ctx := context.Background()
//	payload := &proto.GamerData_MsgQueExit{}
//	err := json.Unmarshal(msg.Payload, payload)
//	if err != nil {
//		log.Error(ctx, "CallbackTest payload:%v, err:%v", payload, err)
//		return false
//	}
//
//	if p.serverId != payload.ServerId { //不是指定伺服器的玩家直接忽略
//		return false
//	}
//	log.Warn(ctx, "OnGamerMsgQueExit gamer:%v serverId:%d", payload.GamerId, payload.ServerId)
//	p.RemoveContainerGamer(payload.GamerId)
//	return true
//}

func (p *GamerContainer) GetGamerCount() int {
	p.mapLock.RLock()
	defer p.mapLock.RUnlock()
	return len(p.GamerMap)
}
