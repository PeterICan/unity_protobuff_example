# 🚀 應用層通訊抽象與協定實驗專案

本專案旨在實作一套 **資料模型 (Data Model)** 與 **通訊協定** 徹底解耦的 Server/Client 架構。我們的核心目標是建立一個統一的業務邏輯層，並在底層實驗兩種截然不同的通訊模式：**高效能的 Protobuf/TCP** 與 **高相容性的 REST API/JSON**。

| 項目 | 技術棧 | 說明 |
| :--- | :--- | :--- |
| **Server** | **Go (Golang)** | 負責統一的業務邏輯層，可啟動 TCP Socket 或 HTTP 伺服器。 |
| **Client** | **Unity (C#)** | 負責統一的業務邏輯層，可切換使用 Socket 或 HTTP 客戶端。 |
| **核心抽象** | **資料模型層** & **傳輸協定層** | 確保業務邏輯不依賴於 Protobuf 或 JSON，也不依賴於 TCP 或 HTTP。 |

---

## 🧪 實驗模式概覽

| 模式 | 傳輸協定 | 資料格式 | 適用場景 |
| :--- | :--- | :--- | :--- |
| **實驗模式 I** | **TCP** (Socket 連線) | **Protocol Buffers** (二進位) | 遊戲內即時數據交換、高頻率同步。 |
| **實驗模式 II** | **HTTP** (REST API) | **JSON** (文字) | 登入、配置獲取、後台管理、低頻率操作。 |

---

## 🛠️ 先決條件

在啟動專案前，請確保您的開發環境已安裝以下工具：

1.  **Go Runtime & SDK**
2.  **Unity Editor (建議版本：LTS)**
3.  **Protocol Buffers 編譯器 (`protoc`)**
    * Go 插件安裝：`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
4.  **NJsonSchema/OpenAPI 相關工具** (用於模式 II 的 C# 客戶端生成與驗證)

---

## 💡 實作待辦事項 (TODO)

為了實現雙重解耦的架構，以下是接下來需要完成的關鍵任務：

1.  **定義統一資料模型 (Data Models):**
    * 在共用層次定義業務所需的 Go Structs / C# Classes (例如 `UserCredentials`, `PlayerPosition`)，**不包含**任何 Protobuf 或 JSON 標籤。
2.  **實作 Protobuf 格式層 (模式 I):**
    * 撰寫 `proto/message.proto`，將步驟 1 的資料模型轉換為 Protobuf 格式。
    * 執行 `protoc` 指令生成 Go 和 C# 程式碼。
3.  **實作 JSON 格式層 (模式 II):**
    * 在 Server 端定義 REST API 路由，使用標準 Go JSON 處理或 NJsonSchema 相關工具進行序列化。
4.  **抽象傳輸層 (IConnection):**
    * 定義 Go/C# 的 `IConnection` 介面，用於發送/接收序列化後的位元組/資料。
5.  **具體傳輸實作:**
    * 實作 `TCPTransport` (用於模式 I)。
    * 實作 `HTTPClientTransport` / `HTTPServer` (用於模式 II)。

---

## ▶️ 快速啟動指南

Server 端請參考 **`Server/README.md`**，Client 端請參考 **`Client/README.md`**，以確定當前要測試的模式及啟動方式。