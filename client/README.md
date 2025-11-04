# 🎮 Unity Client 說明

本專案實現**雙重解耦**的網路通訊架構，支援兩種通訊模式：

---

## 🚀 快速開始

### 1. 設定連線

在 `NetworkManager` 中配置伺服器位址：

- **TCP模式**: `127.0.0.1:8080`
- **WebSocket模式**: `ws://127.0.0.1:8081`

### 2. 選擇通訊模式

修改相關組件中的實例化代碼：

```csharp
// 模式 I: TCP + Protobuf
_connection = new TCPClientTransport();
_serializer = new ProtobufSerializer();

// 模式 II: WebSocket + JSON  
_connection = new WebSocketClientTransport();
_serializer = new JSONSerializer();
```

### 3. 運行測試

在 Unity Editor 中運行主場景，使用 UI 介面測試位置更新和玩家資訊功能。

---

## 📦 建置執行檔

執行根目錄下的建置腳本：

```bash
client\build_project.bat
```

建置完成後，執行檔位於 `Client/Builds/` 目錄。

---

## 📁 專案結構

```
Client/Assets/Scripts/
├── Game/
│   ├── Models/                    # 資料和業務邏輯
│   │   └── ConnectionModel.cs
│   ├── Presenters/                # MVP模式的協調層
│   │   └── ConnectionPresenter.cs
│   ├── Singleton/                 # 全域管理器
│   │   └── SystemManager.cs
│   ├── TestScript/                # 網路功能測試
│   │   ├── AntnetEchoTest.cs
│   │   └── PositionUpdateTest.cs
│   └── Views/                     # UI介面組件
│       ├── ConnectionView.cs
│       ├── IConnectionView.cs
│       └── TabManagerView.cs
└── Network/
    ├── Serializers/               # 序列化抽象層
    │   ├── ISerializer.cs         # 序列化介面
    │   ├── JSONSerializer.cs      # JSON序列化實作
    │   └── ProtobufSerializer.cs  # Protobuf序列化實作
    └── Transports/                # 傳輸抽象層
        ├── IConnection.cs         # 連線介面
        ├── TCPClientTransport.cs  # TCP實作
        └── WebSocketClientTransport.cs # WebSocket實作
```

### 核心組件說明

- **Network層**: 實現雙重抽象的核心，分離序列化與傳輸邏輯
- **Game層**: 採用MVP模式，分離UI、業務邏輯與資料
- **TestScript**: 提供各種網路功能的測試用例

---

## 🏗️ 架構設計

### 抽象層介面

- **`IConnection`**: 抽象網路連線（TCP/WebSocket）
- **`ISerializer`**: 抽象序列化（Protobuf/JSON）

### 通訊模式

| 模式 | 協定 | 用途 | 實作類別 |
|------|------|------|----------|
| **模式 I** | Protobuf + TCP | 即時遊戲數據 | `TCPClientTransport` + `ProtobufSerializer` |
| **模式 II** | JSON + WebSocket | Web友好通訊 | `WebSocketClientTransport` + `JSONSerializer` |