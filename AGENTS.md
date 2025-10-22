# AGENTS.md：專案協作指南與實作清單

## 使用者偏好

- 專案相關的任何文件產出、討論、註解都請使用中文。

## 協作流程 (Collaboration Workflow)

- **任務確認**: 在完成一個主要任務（例如 TODO 列表中的一個編號項）後，我會向您報告進度。在繼續下一個任務之前，我需要得到您的確認。請回覆「好的」、「繼續」或提供進一步的指示，以確保我們的協作保持同步。

## 專案狀態與核心目標

| 項目 | 描述 | 狀態 |
| :--- | :--- | :--- |
| **專案名稱** | 應用層通訊抽象與協定實驗專案 | 規劃調整 |
| **目標** | 實作一個**雙重解耦**的 Server/Client 架構，實驗 **Protobuf/TCP** 與 **JSON/WebSocket** 兩種通訊模式。 | 已確認 |
| **核心技術** | Go Server (**antnet 框架**) / Unity Client / Protobuf / JSON | 已確認 |

---

## 核心設計理念：雙重抽象 (Decoupling)

本專案將業務邏輯與通訊細節分離成兩層抽象。**注意：在新的計畫中，Server 端的抽象層由 `antnet` 框架提供，我們的工作是使用其 API。Client 端仍需自行實現抽象層。**

| 抽象層 | 負責功能 | 關鍵介面 (示例) | 替換實作 (模式 I / 模式 II) |
| :--- | :--- | :--- | :--- |
| **資料格式** | 處理原生模型與二進位/文字間的轉換。 | `ISerializer` | ProtobufSerializer / JSONSerializer |
| **傳輸協定** | 處理網路連線、Socket I/O 或 WebSocket 請求。 | `IConnection` | TCPTransport / WebSocketTransport |

---

## 📋 接下來的 TODO 列表 (Next Steps) - *已調整為位置優先的開發順序*

### 階段一：基礎設定 (1-2)

| # | 任務 | 描述 | 預期輸出檔案/狀態 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 1 | **檢查現有 Go Server 狀態** | 檢查 `server/go.mod` 和 `server/main.go` 的當前狀態，確認 antnet 依賴是否正確設定 | 確認現有設定狀態 | ✅ 已完成 |
| 2 | **更新 antnet 依賴設定** | 確保 `server/go.mod` 中包含正確的 antnet 依賴，如需要則執行 `go get` 更新 | `server/go.mod` 更新 | ✅ 已完成 |

### 階段二：位置功能優先 (3-9)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 3 | **定義位置資料模型** | 在 `common/types.go` 和 `common/Types.cs` 中定義 `PlayerPosition`、`PositionUpdate` 等位置相關的核心資料結構 | `common/types.go`, `common/Types.cs` | ✅ 已完成 |
| 4 | **建立位置處理 Handler** | 建立 `server/handlers/position.go`，處理位置更新和廣播的業務邏輯 | `server/handlers/position.go` | ✅ 已完成 |
| 5 | **實作 antnet TCP 伺服器設定** | 更新 `server/main.go`，設定 antnet 監聽 TCP 埠並使用 Protobuf 解析器處理位置訊息 | `server/main.go` 更新 | ✅ 已完成 |
| 6 | **實作 Client 端抽象介面** | 在 Unity Client 中定義 `ISerializer` 和 `IConnection` 介面，建立抽象層基礎 | `Client/Scripts/ISerializer.cs`, `Client/Scripts/IConnection.cs` | ✅ 已完成 |
| 7 | **實作 ProtobufSerializer** | 在 `Client/Scripts/` 中實作 `ProtobufSerializer.cs`，處理 Protobuf 格式序列化 | `Client/Scripts/ProtobufSerializer.cs` | ✅ 已完成 |
| 8 | **實作 TCPClientTransport** | 在 `Client/Scripts/` 中實作 `TCPClientTransport.cs`，使用 TcpClient 連接到 antnet 伺服器 | `Client/Scripts/TCPClientTransport.cs` | ✅ 已完成 |
| 9 | **進行 TCP/Protobuf 位置測試** | 測試 Client 發送位置更新，Server 處理並回應的完整 TCP 流程 | 功能可運作 | ✅ 已完成 |

### 階段三：WebSocket 模式 (10-13)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 10 | **設定 antnet WebSocket 端點** | 更新 `server/main.go`，新增 WebSocket 監聽埠並使用 JSON 解析器 | `server/main.go` 更新 | ✅ 已完成 |
| 11 | **實作 JSONSerializer** | 在 `Client/Scripts/` 中實作 `JSONSerializer.cs`，處理 JSON 格式序列化 (JSON Schema 檔案位於 `common/json_schema/position.schema.json`；Client 端將使用 `NJsonSchema` 從 JSON Schema 生成 C# 類別；Go 後端將使用 `omissis/go-jsonschema` 從 JSON Schema 生成結構體) | `Client/Scripts/JSONSerializer.cs` | ✅ 已完成 |
| 12 | **實作 WebSocketClientTransport** | 在 `Client/Scripts/` 中實作 `WebSocketClientTransport.cs`，連接到 antnet 的 WebSocket 端點 | `Client/Scripts/WebSocketClientTransport.cs` | ✅ 已完成 |
| 13 | **進行 WebSocket/JSON 位置測試** | 測試 Client 發送位置更新，Server 處理並廣播的完整 WebSocket 流程 | 功能可運作 | ✅ 已完成 |

### 階段四：通訊指令碼定義 (14-15)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 14 | **為 WebSocket 定義獨立指令碼** | **最終設計方案：**

**核心目標**：
實現一個**雙重解耦**的 JSON/WebSocket 通訊模式，使其與 Protobuf/TCP 模式完全獨立，並具備 Web 友好的特性，同時解決 `antnet` 框架內建 JSON 解析器的性能瓶頸。

**設計方案概述**：

1. **通訊協定基礎 (JSON)**：
    - 所有 WebSocket 訊息均採用 **JSON 格式**。
    - 每個 C2S (Client-to-Server) 和 S2C (Server-to-Client) 訊息都包含一個**`route` 欄位 (字串)**，用於明確指示訊息的類型和預期處理邏輯（例如 `"position/update"`）。這取代了舊 TCP 模式中的 `Cmd`/`Act` 整數組合。
    - 訊息中包含一個 `request_id` 欄位 (字串)，用於非同步請求與回應的關聯，但其功能實作已延後。
    - 定義了標準的 `ErrorResponse` 結構，用於統一的錯誤回報。

2. **Protobuf 檔案結構與生成**：
    - `.proto` 檔案作為 JSON 訊息的**結構定義語言**，確保前後端資料結構的一致性。
    - **檔案組織**：`.proto` 檔案依照邏輯套件（`json_api`）分層存放於 `common/proto/json_api/` 目錄下，例如 `common.api.proto`、`position.api.proto`、`gamer_info.api.proto`。
    - **`go_package` 設定**：`.proto` 檔案中的 `option go_package = "/.;json_api";` 確保 Go 程式碼能正確生成到 `server/generated/json_api` 子目錄。
    - **`import` 機制**：通用訊息（如 `ErrorResponse`）定義在 `common.api.proto` 中，並被其他 `.proto` 檔案導入使用。
    - **腳本調整**：`generate_go_proto.bat` 和 `generate_csharp_proto.bat` 腳本已修改，使用 `pushd/popd` 和 `protoc` 兩次呼叫（分別處理 `message.proto` 和 `json_api` 相關 `.proto`），確保所有 `.proto` 檔案都能正確生成對應的 Go 和 C# 程式碼。

3. **客戶端實作 (Unity C#)**：
    - **`JSONSerializer` 增強**：
        - 繼續使用 `Google.Protobuf.JsonFormatter` 和 `JsonParser` 處理 Protobuf 生成的 C# 類型。
        - 引入 `Newtonsoft.Json` 套件，並在 `JSONSerializer` 中新增 `DeserializeToJObject(byte[] data)` 方法，用於將原始 JSON 字串解析為通用的 `JObject`。
    - **兩階段反序列化**：
        1. `OnMessageReceived` 函式首先使用 `DeserializeToJObject` 獲取 `JObject`，從中提取 `route` 字串。
        2. 根據 `route` 的值，動態判斷目標 `S2C` 訊息類型，然後使用 `_jsonSerializer.Deserialize<T>(byte[] data)` 將原始 JSON 完整反序列化為該特定類型。
    - **發送訊息**：C# `C2S` 訊息物件（例如 `C2SPositionUpdate`）被建立，其 `route` 欄位被設定，然後透過 `_jsonSerializer.Serialize` 序列化為 JSON 字節發送。

4. **伺服器端實作 (Go `antnet`)**：
    - **自定義解析器 (`JsonRouteParser`)**：
        - 實作 `server/parser/json_route_parser.go`，它實現了 `antnet.IParser` 介面。
        - `JsonRouteParser.ParseC2S` 方法會先從 incoming JSON 中提取 `route`，然後透過內部 `routeMap` 直接查找並反序列化為正確的 C2S 訊息類型，實現 `O(1)` 的解析效率。
        - `JsonRouteParser.PackMsg` 負責將 Go 物件序列化為 JSON 字節。
    - **`antnet` 框架整合**：
        - **修改 `antnet` 原始碼**：在 `server/third-party/antnet/parser.go` 中為 `antnet.Parser` 結構體添加了公共的 `SetIParser(p antnet.IParser)` 方法，以允許從外部設定其內部的 `parser` 欄位。
        - `server/ws_server.go` 在 `NewWebSocketServer` 中，使用 `baseParser.SetIParser(jsonRouteParser)` 將自定義解析器注入 `antnet` 框架。
        - `JsonRouteParser.RegisterMsg(route string, c2s interface{}, s2c interface{})` 方法用於將 `route` 字串與 C2S/S2C 訊息類型關聯起來，填充 `JsonRouteParser` 內部的 `routeMap`。
    - **回應處理**：處理函式（例如 `HandleC2SPositionUpdate`）手動使用 `encoding/json.Marshal` 將 S2C 回應物件序列化為 JSON 字節，然後放入 `antnet.Message` 的 `Data` 欄位中發送。

**考量**：
- **性能**：自定義 `JsonRouteParser` 顯著提升了訊息分發的性能。
- **第三方修改**：對 `antnet` 原始碼進行了少量修改以實現整合，這是為了解決性能問題而做出的權衡。 | `common/types.go`, `common/Types.cs` 更新 | ✅ 已完成 |

| 15 | **實作自定義 JSON 路由解析器** | 為了避免 `antnet` 內建 `JsonParser` 的「型別猜測」性能問題，實作 `server/parser/json_route_parser.go`，並修改 `antnet` 原始碼以支援其整合。 | `server/parser/json_route_parser.go`，`server/third-party/antnet/parser.go` 修改，`server/ws_server.go` 更新。 | ✅ 已完成 |

### 階段四點五：發布準備 (15.1-15.3)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 15.1 | **生成 Unity Client 執行檔** | 建置 Unity Client 專案，生成可執行的應用程式。 | `Client/Builds/` 下的執行檔 | ✅ 已完成 |
| 15.2 | **生成 Go Server 執行檔** | 編譯 Go Server 專案，生成可執行的伺服器程式。已建立 `server/build_server.bat` 腳本來自動化此過程。 | `server/server_executable` | ✅ 已完成 |
| 15.3 | **建立 Docker Compose 配置** | 編寫 `docker-compose.yml` 文件，用於部署和管理 Server。 | `docker-compose.yml` | ⏹️ 未開始 |

### 階段五：基礎 UI (16)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 16 | **實作基礎 Viewer 介面** | 建立簡化的 `NetworkViewer.cs`，專注於位置更新功能的 UI 介面測試 | `Client/Scripts/NetworkViewer.cs` | ⏹️ 未開始 |

### 階段六：登入功能 (17-19) - *延後執行*

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 17 | **定義登入資料模型** | 在 common 中定義 `UserCredentials`、`LoginResponse` 等登入相關資料結構 | 更新 `common/types.go`, `common/Types.cs` | ⏹️ 未開始 |
| 18 | **建立登入處理 Handler** | 建立 `server/handlers/login.go`，處理登入驗證的業務邏輯 | `server/handlers/login.go` | ⏹️ 未開始 |
| 19 | **完整 Unity Client Viewer 介面** | 擴展 NetworkViewer，新增登入驗證功能和完整的 UI 測試介面 | `Client/Scripts/NetworkViewer.cs` 更新 | ⏹️ 未開始 |

### 階段七：最終整合 (20)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 20 | **整合測試和最終驗證** | 進行完整的雙重抽象架構測試，確認所有功能（位置+登入）正常運作 | 功能可運作 | ⏹️ 未開始 |

### 階段八：實作 request_id (21)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 21 | **實作 request_id 的生成與處理** | 在客戶端生成唯一的 `request_id` 並包含在 C2S 訊息中，伺服器將其回傳於 S2C 訊息。客戶端利用 `request_id` 關聯請求與回應。此功能將在基礎功能穩定後實作。 | Client/Server 程式碼更新，`request_id` 欄位啟用。 | ⏹️ 未開始 |

### 階段九：架構重構 (22)

| # | 任務 | 描述 | 預期輸出檔案 | 狀態 |
| :--- | :--- | :--- | :--- | :--- |
| 22 | **JSON 序列化流程重構** | 修正 JSON 序列化流程，使其使用獨立的資料模型 (POCO 或從 JSON Schema 生成)，以符合「雙重解耦」的架構目標，解決目前對 Protobuf 生成類別的依賴問題。(時間允許下才會進行) | `Client/Scripts/JSONSerializer.cs` 更新 | ⏹️ 未開始 |

---

## 應用提案：網路狀態與互動演示 (Network Status & Interaction Demo)

**1. 應用類型**：Unity UI 應用

**2. 核心功能**：

*   **連線狀態顯示**：
    *   一個文字標籤，實時顯示與伺服器的連線狀態（例如：「已連線」、「未連線」、「連線中...」）。
*   **位置更新功能**：
    *   一個「發送位置更新」按鈕。點擊後，客戶端會生成一組隨機的 X, Y, Z 座標，並透過 `C2SPositionUpdate` 訊息發送給伺服器。
    *   一個文字標籤，顯示最後一次發送的位置數據。
    *   一個文字標籤，顯示從伺服器接收到的 `S2CPositionUpdate` 回應狀態（例如：「位置更新成功」）。
*   **玩家資訊查詢功能**：
    *   一個「查詢玩家資訊」按鈕。點擊後，客戶端會發送 `C2SGamerInfoRetrieve` 訊息給伺服器。
    *   多個文字標籤，顯示從伺服器接收到的 `S2CGamerInfoRetrieve` 訊息中的玩家資訊（例如：暱稱、等級、金錢、寶石）。
*   **日誌輸出面板**：
    *   一個可滾動的文字區域，顯示所有發送和接收到的網路訊息（JSON 格式），以及重要的連線事件日誌。

**3. 互動方式**：
*   所有功能都透過 UI 上的按鈕點擊來觸發。

**4. 視覺風格**：
*   使用 Unity UI (UGUI) 的基礎元件，例如 `Button`、`Text`、`Scroll Rect`。
*   整體介面簡潔，幾片面板即可，無需複雜的視覺設計。
