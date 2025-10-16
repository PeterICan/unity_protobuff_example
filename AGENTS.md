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

| # | 任務 | 描述 | 預期輸出檔案/狀態 |
| :--- | :--- | :--- | :--- |
| 1 | **檢查現有 Go Server 狀態** | 檢查 `server/go.mod` 和 `server/main.go` 的當前狀態，確認 antnet 依賴是否正確設定 | 確認現有設定狀態 |
| 2 | **更新 antnet 依賴設定** | 確保 `server/go.mod` 中包含正確的 antnet 依賴，如需要則執行 `go get` 更新 | `server/go.mod` 更新 |

### 階段二：位置功能優先 (3-9)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 3 | **定義位置資料模型** | 在 `common/types.go` 和 `common/Types.cs` 中定義 `PlayerPosition`、`PositionUpdate` 等位置相關的核心資料結構 | `common/types.go`, `common/Types.cs` |
| 4 | **建立位置處理 Handler** | 建立 `server/handlers/position.go`，處理位置更新和廣播的業務邏輯 | `server/handlers/position.go` |
| 5 | **實作 antnet TCP 伺服器設定** | 更新 `server/main.go`，設定 antnet 監聽 TCP 埠並使用 Protobuf 解析器處理位置訊息 | `server/main.go` 更新 |
| 6 | **實作 Client 端抽象介面** | 在 Unity Client 中定義 `ISerializer` 和 `IConnection` 介面，建立抽象層基礎 | `Client/Scripts/ISerializer.cs`, `Client/Scripts/IConnection.cs` |
| 7 | **實作 ProtobufSerializer** | 在 `Client/Scripts/` 中實作 `ProtobufSerializer.cs`，處理 Protobuf 格式序列化 | `Client/Scripts/ProtobufSerializer.cs` |
| 8 | **實作 TCPClientTransport** | 在 `Client/Scripts/` 中實作 `TCPClientTransport.cs`，使用 TcpClient 連接到 antnet 伺服器 | `Client/Scripts/TCPClientTransport.cs` |
| 9 | **進行 TCP/Protobuf 位置測試** | 測試 Client 發送位置更新，Server 處理並回應的完整 TCP 流程 | 功能可運作 |

### 階段三：WebSocket 模式 (10-13)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 10 | **設定 antnet WebSocket 端點** | 更新 `server/main.go`，新增 WebSocket 監聽埠並使用 JSON 解析器 | `server/main.go` 更新 |
| 11 | **實作 JSONSerializer** | 在 `Client/Scripts/` 中實作 `JSONSerializer.cs`，處理 JSON 格式序列化 | `Client/Scripts/JSONSerializer.cs` |
| 12 | **實作 WebSocketClientTransport** | 在 `Client/Scripts/` 中實作 `WebSocketClientTransport.cs`，連接到 antnet 的 WebSocket 端點 | `Client/Scripts/WebSocketClientTransport.cs` |
| 13 | **進行 WebSocket/JSON 位置測試** | 測試 Client 發送位置更新，Server 處理並廣播的完整 WebSocket 流程 | 功能可運作 |

### 階段四：基礎 UI (14)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 14 | **實作基礎 Viewer 介面** | 建立簡化的 `NetworkViewer.cs`，專注於位置更新功能的 UI 介面測試 | `Client/Scripts/NetworkViewer.cs` |

### 階段五：登入功能 (15-17) - *延後執行*

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 15 | **定義登入資料模型** | 在 common 中定義 `UserCredentials`、`LoginResponse` 等登入相關資料結構 | 更新 `common/types.go`, `common/Types.cs` |
| 16 | **建立登入處理 Handler** | 建立 `server/handlers/login.go`，處理登入驗證的業務邏輯 | `server/handlers/login.go` |
| 17 | **完整 Unity Client Viewer 介面** | 擴展 NetworkViewer，新增登入驗證功能和完整的 UI 測試介面 | `Client/Scripts/NetworkViewer.cs` 更新 |

### 階段六：最終整合 (18)

| # | 任務 | 描述 | 預期輸出檔案 |
| :--- | :--- | :--- | :--- |
| 18 | **整合測試和最終驗證** | 進行完整的雙重抽象架構測試，確認所有功能（位置+登入）正常運作 | 功能可運作 |
