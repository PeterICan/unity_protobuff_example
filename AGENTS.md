# AGENTS.md：專案協作指南與實作清單

## 使用者偏好
討論要用中文

## 專案狀態與核心目標

| 項目 | 描述 | 狀態 |
| :--- | :--- | :--- |
| **專案名稱** | 應用層通訊抽象與協定實驗專案 | 規劃完成 |
| **目標** | 實作一個**雙重解耦**的 Server/Client 架構，實驗 **Protobuf/TCP** 與 **JSON/REST API** 兩種通訊模式。 | 已確認 |
| **核心技術** | Go Server / Unity Client / Protobuf / JSON (NJsonSchema) / 抽象層設計 | 已確認 |

---

## 核心設計理念：雙重抽象 (Decoupling)

本專案將業務邏輯與通訊細節分離成兩層抽象，確保任一格式或協定變動時，業務程式碼無需修改。

| 抽象層 | 負責功能 | 關鍵介面 (示例) | 替換實作 (模式 I / 模式 II) |
| :--- | :--- | :--- | :--- |
| **資料格式** | 處理原生模型與二進位/文字間的轉換。 | `ISerializer` | ProtobufSerializer / JSONSerializer |
| **傳輸協定** | 處理網路連線、Socket I/O 或 HTTP 請求。 | `IConnection` | TCPTransport / HttpClientTransport |

---

## 📋 接下來的 TODO 列表 (Next Steps)

以下是從規劃階段進入實作階段的關鍵步驟，請按順序完成：

### 階段一：模型與格式基礎 (Common Base & Protobuf Definition)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 1 | **定義統一資料模型 (Data Models)** | 定義核心業務 Go Structs / C# Classes，**不包含**任何 Protobuf 或 JSON 標籤。 | `common/types.go`, `common/Types.cs` |
| 2 | **Protobuf 檔案定義 (模式 I)** | 撰寫 `message.proto`，包含 `UserCredentials`, `PlayerPosition` 等，並定義 **`WrapperMessage`** 以統一傳輸。 | `proto/message.proto` |
| 3 | **生成 Protobuf 程式碼** | 執行 `protoc` 指令，生成 Go 和 C# 語言的 Protobuf 程式碼。 | `Server/generated/*.go`, `Client/Assets/Generated/*.cs` |

### 階段二：序列化層實作 (Serializer Implementation)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 4 | **定義 `ISerializer` 介面** | 定義通用的序列化/反序列化介面。 | `common/ISerializer.go`, `common/ISerializer.cs` |
| 5 | **實作 `ProtobufSerializer`** | 實作 `ISerializer` 介面，處理 Protobuf 格式的序列化、反序列化及 `WrapperMessage` 封裝。 | `Server/pkg/ProtobufSerializer.go`, `Client/Scripts/ProtobufSerializer.cs` |
| 6 | **實作 `JSONSerializer`** | 實作 `ISerializer` 介面，處理標準 JSON 格式的轉換。 | `Server/pkg/JSONSerializer.go`, `Client/Scripts/JSONSerializer.cs` |

### 階段三：TCP 傳輸層實作 (模式 I：Socket 連線)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 7 | **定義 `IConnection` 介面** | 定義抽象的連線、發送 (`Send(data []byte)`) 和接收事件/方法。 | `common/IConnection.go`, `common/IConnection.cs` |
| 8 | **實作 `TCPTransport` (Go Server)** | 實作 `IConnection` 介面，負責 `net.Listener` 和 **Length-Prefixing** 的讀取邏輯。 | `Server/pkg/TCPTransport.go` |
| 9 | **實作 `TCPClientTransport` (Unity Client)** | 實作 `IConnection` 介面，負責 `TcpClient`、非同步連線和 **Length-Prefixing** 的讀寫邏輯。 | `Client/Scripts/TCPClientTransport.cs` |

### 階段四：REST API 傳輸層實作 (模式 II：HTTP)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 10 | **實作 `HTTPServer` (Go Server)** | 設置 Go 的 `net/http` 路由，並使用 `JSONSerializer` 處理請求和回應。 | `Server/pkg/HTTPServer.go` |
| 11 | **實作 `HttpClientTransport` (Unity Client)** | 實作 `IConnection` 介面，使用 `UnityWebRequest` 發送 HTTP 請求並使用 `JSONSerializer` 處理 JSON 響應。 | `Client/Scripts/HttpClientTransport.cs` |

---