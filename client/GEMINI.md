# MVP Planning for Gemini

**注意：** 所有討論與註解請使用繁體中文。

---

### 1. 這個 MVP 旨在解決的核心問題是什麼？

驗證並展示一個靈活的網路架構，允許 Unity 客戶端在不修改主要遊戲邏輯的情況下，能夠輕鬆切換兩種截然不同的後端通訊協定，並處理**相同的資料**：

1. **高效能模式 (Protobuf + TCP):** 用於處理即時、低延遲的遊戲數據。
2. **高相容性模式 (JSON + HTTP):** 用於處理登入、身份驗證等非即時性操作。

這個 MVP 的成功標準是證明此架構的可行性與實用性，為未來快速開發和迭代打下基礎，並著重於**實驗兩種通訊方式在處理相同資料時的表現與切換機制**。

---

### 2. 目標受眾是誰？

內部開發後端人員與前端人員，用於研究目的。

---

### 3. 為了解決核心問題，目標受眾所需的最低限度功能有哪些？

**應用提案：網路狀態與互動演示 (Network Status & Interaction Demo)**

1. **應用類型：** Unity UI 應用

2. **核心功能 (頁籤式佈局):**

    * **頁籤一：連線狀態**
        * **連線狀態顯示：** 一個文字標籤，實時顯示與伺服器的連線狀態（例如：「已連線」、「未連線」、「連線中...」）。
        * **其他玩家狀態顯示：** 一個專門的 UI 區域，用來顯示從伺服器廣播的其他玩家的連線狀態和即時位置。
        * **日誌輸出面板：** 一個可滾動的文字區域，顯示所有發送和接收到的網路訊息（JSON 格式），以及重要的連線事件日誌。

    * **頁籤二：測試功能**
        * **其他玩家狀態顯示：** 一個專門的 UI 區域，用來顯示從伺服器廣播的其他玩家的連線狀態和即時位置。
        * **位置更新功能：**
            * 一個「發送位置更新」按鈕。點擊後，客戶端會生成一組隨機的 X, Y, Z 座標，並透過 `C2SPositionUpdate` 訊息發送給伺服器。
            * 一個文字標籤，顯示最後一次發送的位置數據。
            * 一個文字標籤，顯示從伺服器接收到的 `S2CPositionUpdate` 回應狀態（例如：「位置更新成功」）。

    * **頁籤三：玩家資訊**
        * **玩家資訊查詢功能：**
            * 一個「查詢玩家資訊」按鈕。點擊後，客戶端會發送 `C2SGamerInfoRetrieve` 訊息給伺服器。
            * 多個文字標籤，顯示從伺服器接收到的 `S2CGamerInfoRetrieve` 訊息中的玩家資訊（例如：暱稱、等級、金錢、寶石）。

3. **互動方式：**
    * 所有功能都透過 UI 上的按鈕點擊來觸發。

4. **視覺風格：**
    * 使用 Unity UI (UGUI) 的基礎元件，例如 Button、Text、Scroll Rect，並包含一個頁籤系統 (Tab System) 來組織介面。
    * 整體介面簡潔，幾片面板即可，無需複雜的視覺設計。

---

### 4. 是否有任何特定的技術或限制需要注意？

* **必須**使用現有的 `IConnection` 和 `ISerializer` 介面抽象。
* **必須**實現 `TCPClientTransport` (Protobuf) 和 `WebSocketClientTransport` (JSON) 兩種模式的切換能力。
* **必須**使用 Unity UI (UGUI) 進行介面開發。
* 客戶端生成的 C# 資料模型需與 Protobuf 和 JSON 格式相容。
* **重要：** 前端傳往後端的 JSON 訊息，其內容是由 Protobuf 序列化而來。同時，收到的訊息也會依據封包上的路由 (route) 進行反序列化。這表示即使是 JSON 模式，底層的資料結構和處理邏輯仍基於 Protobuf 定義。

---

### 5. 階段性 TODO 清單 (基於 MVP 架構)

#### Phase 1: 基礎架構與核心元件 (Foundation & Core Components)

* [x] **設定 Unity 專案與 MVP 架構基礎:**
  * 建立必要的資料夾結構 (e.g., Models, Views, Presenters)。
  * 確認專案設定符合開發需求。
  * 狀態：完成
* [x] **確認 `IConnection` 和 `ISerializer` 介面已存在:**
  * 檢查現有介面定義是否符合需求。
  * 狀態：完成
* [x] **實作 `TCPClientTransport` 和 `WebSocketClientTransport`:**
  * 確保兩種傳輸層實作符合 `IConnection` 介面。
  * 狀態：完成
* [x] **實作 `ProtobufSerializer` 和 `JSONSerializer`:**
  * 確保兩種序列化器實作符合 `ISerializer` 介面。
  * 狀態：完成
* [x] **建立基礎的資料模型 (C# Classes) 與 Protobuf 定義:**
  * 根據 `C2SPositionUpdate`, `S2CPositionUpdate`, `C2SGamerInfoRetrieve`, `S2CGamerInfoRetrieve` 等訊息定義 C# 資料模型。
  * 狀態：完成

#### Phase 2: UI 介面設計與實作 (UI Design & Implementation)

* [x] **設計主 UI 佈局，包含頁籤系統 (Tab System):**
  * 在 Unity 中建立 Canvas 和基礎 UI 元件。
  * 實作頁籤切換邏輯 (使用 `TabManagerView.cs`)。
  * 狀態：完成
* [x] **實作「連線狀態」頁籤的 UI 元件:**
  * 連線狀態文字 (Text)。
  * 其他玩家狀態 UI 區域 (e.g., ScrollRect + Text/Prefab)。
  * 日誌輸出面板 (ScrollRect + Text)。
  * 狀態：完成
* [x] **實作「測試功能」頁籤的 UI 元件:**
  * 「發送位置更新」按鈕 (Button)。
  * 最後一次發送的位置數據顯示 (Text)。
  * `S2CPositionUpdate` 回應狀態顯示 (Text)。
  * 狀態：完成
* [x] **實作「玩家資訊」頁籤的 UI 元件:**
  * 「查詢玩家資訊」按鈕 (Button)。
  * 玩家資訊 (多個 Text for Nickname, Level, Money, Gem)。
  * 狀態：完成
* [ ] **實作「玩家資訊」頁籤的 UI 元件:**
  * 「查詢玩家資訊」按鈕 (Button)。
  * 玩家資訊 (多個 Text for Nickname, Level, Money, Gem)。
  * 狀態：待辦

#### Phase 3: 功能實作與 MVP 邏輯 (Feature Implementation & MVP Logic)

* [x] **連線管理 Presenter/Model:**
  * 實作 `ConnectionModel` (位於 `Assets/Scripts/Game/Models/ConnectionModel.cs`) 處理連線/斷線狀態。
  * 實作 `ConnectionPresenter` (位於 `Assets/Scripts/Game/Presenters/ConnectionPresenter.cs`) 負責 View 與 Model 之間的互動，更新 `ConnectionView` (位於 `Assets/Scripts/Game/Views/ConnectionView.cs`)。
  * 狀態：完成
* [ ] **位置更新 Presenter/Model:**
  * 實作 `PositionUpdateModel` 處理隨機座標生成、發送 `C2SPositionUpdate`。
  * 實作 `PositionUpdatePresenter` 處理按鈕事件、接收 `S2CPositionUpdate` 並更新 `PositionUpdateView`。
  * 狀態：待辦
* [ ] **玩家資訊查詢 Presenter/Model:**
  * 實作 `GamerInfoModel` 處理發送 `C2SGamerInfoRetrieve`。
  * 實作 `GamerInfoPresenter` 處理按鈕事件、接收 `S2CGamerInfoRetrieve` 並更新 `GamerInfoView`。
  * 狀態：待辦
* [ ] **其他玩家狀態 Presenter/Model:**
  * 實作 `OtherPlayersModel` 處理接收廣播訊息。
  * 實作 `OtherPlayersPresenter` 負責更新 `OtherPlayersView`。
  * 狀態：待辦
* [ ] **日誌輸出 Presenter/Model:**
  * 實作 `LogModel` 收集網路訊息和事件。
  * 實作 `LogPresenter` 負責更新 `LogView`。
  * 狀態：待辦

#### Phase 4: 整合與測試 (Integration & Testing)

* [ ] **整合所有 UI、Presenter 和 Model:**
  * 確保各個 MVP 元件之間正確連接和協同工作。
  * 狀態：待辦
* [ ] **實作兩種通訊模式 (TCP/Protobuf, HTTP/JSON) 的切換機制:**
  * 設計一個配置介面或邏輯，允許在運行時切換 `IConnection` 和 `ISerializer` 的實作。
  * 狀態：待辦
* [ ] **進行功能測試:**
  * 驗證所有功能在 TCP/Protobuf 模式下均正常運作。
  * 驗證所有功能在 HTTP/JSON 模式下均正常運作。
  * 狀態：待辦
