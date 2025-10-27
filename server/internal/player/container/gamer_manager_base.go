package container

import (
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/internal/proto/handlers"
	"proto_buffer_example/server/internal/subsystem/base"
	"proto_buffer_example/server/third-party/antnet"
	"proto_buffer_example/server/tools/customize"

	"sync"
)

// SubsystemContainer 子系統容器與管理模組
type SubsystemContainer struct {
	ownLock        sync.RWMutex
	ControllerList []base.ISubsystemController
	ModelList      []base.ISubsystemModel
	ServerId       int32
}

//func (p *SubsystemContainer) initAttachedSubsystem(serviceId int32) {
//	for i := range p.ModelList {
//		p.ModelList[i].Init(serviceId, serviceInfo)
//	}
//
//	for i := range p.ControllerList {
//		p.ControllerList[i].Init(serviceId)
//	}
//
//	for i := range p.ModelList {
//		p.ModelList[i].LoadFromDB()
//	}
//}

func (p *SubsystemContainer) InitServerBase(serverId int32, addr string) { //serverInfo server.ServerInfo, port int) {
	p.ServerId = serverId
	//ctx := context.Background()
	//ctx = log.WriteServerIdWithCtx(ctx, p.ServerId)
	//p.initAttachedSubsystem(p.ServerId, serviceInfo)
	msgHandler := &customize.MsgHandler{}
	msgParser := &antnet.Parser{Type: antnet.ParserTypePB}
	p.initRegisterMsgHandler(msgHandler)
	p.initRegisterMsgParser(msgParser, nil)

	//addr := fmt.Sprintf("tcp://:%s", serviceInfo.Port)
	err := antnet.StartServer(addr, antnet.MsgTypeMsg, msgHandler, msgParser)
	if err != nil {
		//log.Error(ctx, "InitServerBase err:%v", err)
		panic(err)
	}
	//log.Info(ctx, "InitServerBase ServerId:%v addr:%v", p.ServerId, addr)
	p.onStart()
	//p.serverMonitor()
}

func (p *SubsystemContainer) InitWebSocketServerBase(serverId int32, addr string) { //serverInfo server.ServerInfo, port int) {
	p.ServerId = serverId
	//ctx := context.Background()
	//ctx = log.WriteServerIdWithCtx(ctx, p.ServerId)
	//p.initAttachedSubsystem(p.ServerId, serviceInfo)
	msgHandler := &customize.MsgHandler{}
	positionHandler := &handlers.PositionHandler{}
	gamerInfoHandler := &handlers.GamerInfoHandler{}
	// Register handlers for the C2S message types
	msgHandler.RegisterMsg(&json_api.C2SPositionUpdate{}, positionHandler.HandleC2SPositionUpdate)
	msgHandler.RegisterMsg(&json_api.C2SGamerInfoRetrieve{}, gamerInfoHandler.HandleC2SGamerInfoRetrieve)

	msgParser := &antnet.Parser{}
	msgParser.Type = antnet.ParserTypeCustom
	msgParser.ErrType = antnet.ParseErrTypeSendRemind
	// 使用自訂的JsonRouteParser來解析封包
	// JsonRouteParser會根據封包中的Route來決定使用哪個Proto結構來解析
	// EX: Route = "position/update" -> 解析成 json_api.C2SPositionUpdate
	jsonRouteParser := customize.NewJsonRouteParser(msgParser)
	msgParser.SetIParser(jsonRouteParser)

	//p.initRegisterMsgHandler(msgHandler)
	// 註冊子系統的封包解析，多傳入一個jsonRouteParser
	// 註冊時一樣使用 jsonRouteParser.RegisterMsg 來註冊封包
	// EX: jsonRouteParser.RegisterMsg("position/update", &json_api.C2SPositionUpdate{}, nil)
	p.initRegisterMsgParser(msgParser, jsonRouteParser)
	//addr := fmt.Sprintf("ws://:%s/ws", serviceInfo.Port)
	err := antnet.StartServer(addr, antnet.MsgTypeCmd, msgHandler, msgParser)
	if err != nil {
		//log.Error(ctx, "InitServerBase err:%v", err)
		panic(err)
	}
	//log.Info(ctx, "InitServerBase ServerId:%v addr:%v", p.ServerId, addr)
	p.onStart()
	//p.serverMonitor()
}

func (p *SubsystemContainer) initRegisterMsgHandler(handler *customize.MsgHandler) {
	for i := range p.ControllerList {
		p.ControllerList[i].RegisterMsgHandler(handler)
	}
}

func (p *SubsystemContainer) initRegisterMsgParser(parser *antnet.Parser, customParser *customize.JsonRouteParser) {
	for i := range p.ControllerList {
		p.ControllerList[i].RegisterMsgParser(parser, nil)
	}
}

func (p *SubsystemContainer) saveAllModel() {
	for i := range p.ModelList {
		p.ModelList[i].SaveToDB()
	}
}

func (p *SubsystemContainer) onStart() {
	for i := range p.ModelList {
		p.ModelList[i].Start()
	}
	for i := range p.ControllerList {
		p.ControllerList[i].Start()
	}
}

func (p *SubsystemContainer) onStop() {
	for i := range p.ModelList {
		p.ModelList[i].Stop()
	}
	for i := range p.ControllerList {
		p.ControllerList[i].Stop()
	}
}

func (p *SubsystemContainer) Stop() {
	p.ownLock.Lock()
	defer p.ownLock.Unlock()
	//ctx := context.Background()
	//ctx = log.WriteServerIdWithCtx(ctx, p.ServerId)
	p.saveAllModel()
	p.onStop()

	//log.Warn(ctx, "\n*********************************************************************************\n"+
	//	" 			    SubsystemContainer Server Stop 		  \n"+
	//	"*********************************************************************************\n")
}

//func (p *SubsystemContainer) serverMonitor() {
//	antnetStatus := antnet.GetStatus()
//	systemInfo := systemhelper.GetSystemInfo()
//
//	tags := influxdefine.GetTagWithServerId(p.ServerId)
//	fields := map[string]interface{}{
//		influxdefine.MemoryUsageField.String():        systemInfo.MemoryUsedPercent,
//		influxdefine.CPUUsageField.String():           systemInfo.CPUTotalUsedPercent,
//		influxdefine.MsgqueCountField.String():        antnetStatus.MsgqueCount,
//		influxdefine.PanicCountField.String():         antnetStatus.PanicCount,
//		influxdefine.GoRoutineCountField.String():     antnetStatus.GoCount,
//		influxdefine.PoolGoRoutineCountField.String(): antnetStatus.PoolGoCount,
//	}
//
//	point := influxhelper.GenPoint("SystemInfo", tags, fields)
//	influxhelper.GetInfluxHelper().WritePoint(point)
//	time.AfterFunc(time.Second*30, p.serverMonitor)
//}
