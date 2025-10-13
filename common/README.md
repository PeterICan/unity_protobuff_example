# 共用專案檔案

此目錄包含整個專案（包括 server 和 client）所共用的資源。

## Protobuf 程式碼生成

此目錄存放了從 .proto 定義檔生成特定語言程式碼的腳本。

### 腳本

- `generate_go_proto.bat`: 生成用於 Server 的 Go 程式碼。
    - **來源**: `../proto/message.proto`
    - **輸出**: `../server/generated/message.pb.go`
- `generate_csharp_proto.bat`: 生成用於 Unity Client 的 C# 程式碼。
    - **來源**: `../proto/message.proto`
    - **輸出**: `../client/Assets/Generated/Message.cs`

### 使用方式

當 `../proto/message.proto` 檔案被修改後，請切換到此 (`common`) 目錄並執行相應的腳本以重新生成程式碼。