# AGENTS.md：專案協作指南與實作清單

## 使用者偏好
- 專案相關的任何文件產出、討論、註解都請使用中文。

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

## 📋 接下來的 TODO 列表 (Next Steps) - *已更新為 antnet 整合計畫*

### 階段一：Server 基礎設定 (antnet)

| # | 任務 | 描述 | 預期輸出檔案/狀態 |
| :--- | :--- | :--- | :--- |
| 1 | **初始化 Go Server 專案** | 在 `server` 目錄下建立 Go module，並新增 `main.go` 作為入口。 | `server/go.mod`, `server/main.go` |
| 2 | **引入 `antnet` 依賴** | 編輯 `go.mod` 或使用 `go get` 來下載 `antnet` 函式庫。 | `server/go.mod` 更新 |
| 3 | **建立基本 `antnet` 伺服器** | 在 `main.go` 中撰寫啟動一個 `antnet` 伺服器實例的基礎程式碼。 | `server/main.go` 可執行 |
| 4 | **定義業務邏輯 Handler** | 建立處理 `UserCredentials` 和 `PlayerPosition` 的 `Handler` 檔案。 | `server/handlers/login.go`, `server/handlers/position.go` |

### 階段二：Client 端抽象層實作

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 5 | **定義核心資料模型** | 在 `common/types.go` 和 `common/Types.cs` 中定義 `UserCredentials` 和 `PlayerPosition` 的原生結構。 | `common/types.go`, `common/Types.cs` |
| 6 | **實作 `ProtobufSerializer` (Client)** | 在 Unity Client 中實作 `ISerializer` 介面，處理 Protobuf 格式。 | `Client/Scripts/ProtobufSerializer.cs` |
| 7 | **實作 `JSONSerializer` (Client)** | 在 Unity Client 中實作 `ISerializer` 介面，處理 JSON 格式。 | `Client/Scripts/JSONSerializer.cs` |

### 階段三：模式 I 整合 (Protobuf/TCP)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 8 | **設定 `antnet` TCP 端點** | 設定 `antnet` 伺服器監聽一個 TCP 埠，並使用 Protobuf 解析器。 | `server/main.go` 更新 |
| 9 | **實作 `TCPClientTransport` (Client)** | 在 Unity Client 中實作 `IConnection` 介面，用 `TcpClient` 連接到 `antnet`。 | `Client/Scripts/TCPClientTransport.cs` |
| 10 | **端對端測試 (TCP)** | 實現 Client 發送登入請求，Server 處理並回應的完整流程。 | 功能可運作 |

### 階段四：模式 II 整合 (JSON/WebSocket)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 11 | **設定 `antnet` WebSocket 端點** | 設定 `antnet` 伺服器監聽一個 WebSocket 埠，並使用 JSON 解析器。 | `server/main.go` 更新 |
| 12 | **實作 `WebSocketClientTransport` (Client)** | 在 Unity Client 中實作 `IConnection` 介面，連接到 `antnet` 的 WebSocket 端點。 | `Client/Scripts/WebSocketClientTransport.cs` |
| 13 | **端對端測試 (WebSocket)** | 實現 Client 發送位置更新，Server 處理並廣播的完整流程。 | 功能可運作 |