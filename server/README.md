# 💻 Go Server 說明

本伺服器使用 **Go (Golang)** 實作，專注於維護核心業務邏輯的**中立性**。它透過抽象化層，能夠切換啟動 **TCP Socket 伺服器** (用於 Protobuf) 或 **HTTP 伺服器** (用於 JSON/REST API)。

---

## 核心設計理念：雙重抽象

### 1. 服務層 (Service Layer)
* **目標：** 核心業務邏輯（例如處理登入請求）僅與**統一的 Go Structs** 互動，完全隔離於 Protobuf 或 JSON 格式。
* **關鍵：** 業務邏輯不應包含任何網路 I/O 或序列化/反序列化程式碼。

### 2. 協定層 (Protocol Layer)
* **`IConnection` 介面：** 定義抽象的發送/接收機制，讓上層業務邏輯得以呼叫。
* **`ISerializer` 介面：** 定義抽象的序列化/反序列化介面，具體實作可以是 `ProtobufSerializer` 或 `JSONSerializer`。

---

## 模式 I：高效能 (Protobuf + TCP)

此模式用於處理即時、高頻率的數據交換。

### 協定實作細節

* **Go 網路:** 使用 `net.Listener` 建立 TCP 伺服器。
* **傳輸協定:** 實作 **`TCPTransport`** 類別來具體執行 `IConnection` 介面。
* **封包分隔:** 實作 **長度前綴 (Length-Prefixing)** 機制，格式為 `[4-byte Length] [Protobuf Wrapper Binary Data]`。
* **序列化:** 使用 **`ProtobufSerializer`** 負責將統一的 Go Structs 轉換為 Protobuf 格式，並封裝進 `WrapperMessage` 中。

### ▶️ 啟動步驟 (模式 I)

1.  **進入 Server 目錄：**
    ```bash
    cd Server
    ```
2.  **啟動伺服器：**
    * 假設 `main.go` 中已配置啟動 TCP 模式的邏輯。
    ```bash
    go run main.go --mode=tcp
    ```
    * **預設埠號：8080**。

---

## 模式 II：高相容性 (JSON + REST API)

此模式用於處理非即時、可靠性要求高的配置或管理操作。

### 協定實作細節

* **Go 網路:** 使用 `net/http` 建立 HTTP 伺服器。
* **傳輸協定:** 使用標準 HTTP 請求/回應週期，不使用 `IConnection` 抽象。
* **資料格式:** 使用 **`JSONSerializer`** 負責將收到的 JSON 資料反序列化為統一的 Go Structs，並將回傳的 Go Structs 序列化為 JSON。
* **文件/驗證:** 可利用 **NJsonSchema/OpenAPI** 相關工具來驗證請求與回應的 JSON 格式。

### ▶️ 啟動步驟 (模式 II)

1.  **進入 Server 目錄：**
    ```bash
    cd Server
    ```
2.  **啟動伺服器：**
    * 假設 `main.go` 中已配置啟動 HTTP 模式的邏輯。
    ```bash
    go run main.go --mode=http --port=8081
    ```
    * **預設埠號：8081** (與 TCP 模式區分)。

---
