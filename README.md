# 🚀 應用層通訊抽象與協定實驗專案

本專案旨在實作一套 **資料模型 (Data Model)** 與 **通訊協定** 徹底解耦的 Server/Client 架構。我們的核心目標是建立一個統一的業務邏輯層，並在底層實驗兩種截然不同的通訊模式：**高效能的 Protobuf/TCP** 與 **高相容性的 REST API/JSON**。

| 項目 | 技術棧 | 說明 |
| :--- | :--- | :--- |
| **Server** | **Go (使用 `antnet` 框架)** | 負責統一的業務邏輯層，利用 `antnet` 快速啟動 TCP Socket 或 WebSocket 伺服器。 |
| **Client** | **Unity (C#)** | 負責統一的業務邏輯層，可切換使用 Socket 或 WebSocket 客戶端。 |
| **核心抽象** | **資料模型層** & **傳輸協定層** | 確保業務邏輯不依賴於 Protobuf 或 JSON，也不依賴於 TCP 或 HTTP/WebSocket。 |

---

## 🧪 實驗模式概覽

| 模式 | 傳輸協定 | 資料格式 | 適用場景 |
| :--- | :--- | :--- | :--- |
| **實驗模式 I** | **TCP** (Socket 連線) | **Protocol Buffers** (二進位) | 遊戲內即時數據交換、高頻率同步。 |
| **實驗模式 II** | **WebSocket** | **JSON** (文字) | 登入、配置獲取、後台管理、低頻率操作。 |

---

## 💡 實作待辦事項 (TODO)

在決定使用 `antnet` 框架後，我們的開發任務將聚焦於業務邏輯和客戶端的實現：

1.  **Server: 初始化 `antnet` 專案**:
    *   設定 Go Server 專案並引入 `antnet` 函式庫。
2.  **Server: 定義訊息處理器 (Handlers)**:
    *   在 `antnet` 中實作 `Handler` 來處理 `UserCredentials` 和 `PlayerPosition` 等業務邏輯。
3.  **Client: 建立通訊模組**:
    *   在 Unity 中實作 `TCPClientTransport` 和 `WebSocketClientTransport` 以連接到 `antnet` 伺服器。
4.  **整合測試 (模式 I - Protobuf/TCP)**:
    *   讓 Client 透過 TCP 發送 Protobuf 訊息給 Server，並驗證業務邏輯。
5.  **整合測試 (模式 II - JSON/WebSocket)**:
    *   讓 Client 透過 WebSocket 發送 JSON 訊息給 Server，並驗證業務邏輯。

---

## ▶️ 快速啟動指南

Server 端請參考 **`Server/README.md`**，Client 端請參考 **`Client/README.md`**，以確定當前要測試的模式及啟動方式。