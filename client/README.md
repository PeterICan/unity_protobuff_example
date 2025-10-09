# 🎮 Unity Client 說明

本專案是 Unity 遊戲客戶端，其核心架構設計是為了實現**應用層與傳輸層的雙重解耦**。遊戲邏輯只處理統一的 C# 資料模型，並可透過配置輕鬆切換到 **Protobuf/TCP** 或 **JSON/HTTP** 模式。

---

## 核心設計理念：雙重抽象

### 1. 服務層 (Service Layer)

* **目標：** 遊戲中的業務邏輯（例如發送移動指令、接收登入結果）僅與**統一的 C# Classes** 互動，完全隔離於 Protobuf 格式或 JSON 格式。
* **關鍵：** 所有的序列化和網路通訊細節都在底層處理。

### 2. 協定層 (Protocol Layer)

* **`IConnection` 介面：** 定義抽象的連線、發送 (`Send`) 和接收事件，讓上層業務邏輯得以呼叫，而無需知道底層是 Socket 還是 HTTP。
* **`ISerializer` 介面：** 定義抽象的序列化/反序列化介面，具體實作可以是 `ProtobufSerializer` 或 `JSONSerializer`。

---

## 模式 I：高效能 (Protobuf + TCP)

此模式用於與 Go Server 建立長連線 (Persistent Connection)，處理即時、低延遲的遊戲數據。

### 協定實作細節

* **C# 網路:** 使用 `System.Net.Sockets.TcpClient` 實現非同步連線。
* **傳輸實作:** **`TCPClientTransport`** 類別實現 `IConnection` 介面。
* **封包處理:** 實作 **長度前綴 (Length-Prefixing)** 機制，從 TCP 位元組流中可靠地讀取完整的 Protobuf 封包。
* **序列化:** **`ProtobufSerializer`** 負責將統一 C# Classes 轉換為 Protobuf 二進位格式。
* **線程安全:** 所有網路讀寫都在後台線程（或 Task）中，透過主線程調度器 (Dispatcher) 將結果安全地回傳給 Unity 主循環。

### ▶️ 運行步驟 (模式 I)

1.  **配置 Server IP/Port：** 在 `NetworkManager` 或配置腳本中，設定正確的 Server IP 和 Port (e.g., `8080`)。
2.  **切換模式：** 確保配置為實例化 `TCPClientTransport`。
3.  **運行：** 在 Unity Editor 中運行主場景。

---

## 模式 II：高相容性 (JSON + REST API)

此模式用於處理登入、配置獲取、或其他無需長連線的業務操作。

### 協定實作細節

* **C# 網路:** 使用 `UnityEngine.Networking.UnityWebRequest` 或 `HttpClient` 處理 HTTP 請求。
* **傳輸實作:** **`HttpClientTransport`** 類別實現 `IConnection` 介面。每個 `Send` 操作將觸發一個獨立的 HTTP 請求。
* **序列化:** **`JSONSerializer`** 負責將統一 C# Classes 序列化為 JSON，並處理從 Server 返回的 JSON 響應。
* **NJsonSchema/OpenAPI:** 可利用生成的 C# 客戶端代碼或模式驗證工具，確保 JSON 資料的格式正確性。

### ▶️ 運行步驟 (模式 II)

1.  **配置 Server URL：** 在 `NetworkManager` 或配置腳本中，設定正確的 REST API URL (e.g., `http://127.0.0.1:8081/api/login`)。
2.  **切換模式：** 確保配置為實例化 `HttpClientTransport`。
3.  **運行：** 在 Unity Editor 中運行主場景。

---
