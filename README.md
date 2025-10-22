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



### Unity Client 建置

Unity Client 的建置過程已透過 `client/build_project.bat` 腳本自動化。此腳本會根據專案設定自動偵測所需的 Unity 編輯器版本，並執行批次模式建置。

**使用方式：**

1.  確保您的系統已安裝 Unity Hub，並且專案所需的 Unity 編輯器版本已透過 Unity Hub 安裝。
2.  開啟命令提示字元或 PowerShell。
3.  導航至專案根目錄。
4.  執行以下命令：
    ```bash
    client\build_project.bat
    ```
5.  建置完成後，可執行檔將位於 `Client/Builds/` 目錄下。

### Go Server 建置

Go Server 的建置過程已透過 `server/build_server.bat` 腳本自動化。

**使用方式：**

1.  開啟命令提示字元或 PowerShell。
2.  導航至專案根目錄。
3.  執行以下命令：
    ```bash
    server\build_server.bat
    ```
4.  建置完成後，可執行檔將位於 `server/server_executable`。

## ▶️ 快速啟動指南

Server 端請參考 **`Server/README.md`**，Client 端請參考 **`Client/README.md`**，以確定當前要測試的模式及啟動方式。