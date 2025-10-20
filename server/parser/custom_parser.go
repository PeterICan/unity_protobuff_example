package parser

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"proto_buffer_example/server/generated"
)

type CustomParser struct{}

func ParseC2S() interface{} {
	// Custom parser implementation goes here
	return &generated.PlayerPosition{}
}

func CustomParserFunc() interface{} {
	// Custom parser implementation goes here
	// 1. 設定解析選項 (UnmarshalOptions)
	unmarshaler := protojson.UnmarshalOptions{}

	// 2. 準備一個空的 Protobuf Message 容器
	newUserMessage := &generated.PlayerPosition{}

	// 3. 執行解析
	jsonBytes := []byte(`{"x":1.0,"y":2.0,"z":3.0}`) // 假設這是從網路接收到的 JSON 資料
	err := unmarshaler.Unmarshal(jsonBytes, newUserMessage)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %v", err)
	}
	fmt.Println("\n--- JSON To Protobuf Message ---")
	// 驗證解析結果
	fmt.Printf("message: %v\n", newUserMessage)

	return newUserMessage
}
