package persistence

import (
	"log"
	"proto_buffer_example/server/tools/common"
)

/*
persistence 套件提供資料持久化相關的功能和工具。
它包含了資料存取和管理的相關程式碼
*/

func GetNextGamerId() int32 {
	storageData := getStorageJson()
	var nextGamerId int32 = 1
	if storageData == nil {
		storageData = make(map[string]interface{})
	}
	if val, ok := storageData[NextGamerIdKey]; ok {
		nextGamerId = int32(val.(float64))
	}
	log.Default().Println("GetNextGamerId:", nextGamerId)
	storageData[NextGamerIdKey] = nextGamerId + 1
	writeStorageJson(storageData)
	return nextGamerId
}

const StorageJsonFile = "data/storage"
const NextGamerIdKey = "NextGamerId"

func getStorageJson() map[string]interface{} {
	//load from StorageJsonFile
	var storageData map[string]interface{}
	common.LoadJsonData(StorageJsonFile, &storageData)
	return storageData
}

func writeStorageJson(storageData map[string]interface{}) {
	//write to StorageJsonFile
	common.WriteJsonData(StorageJsonFile, storageData)
}
