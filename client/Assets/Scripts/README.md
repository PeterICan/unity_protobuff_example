# Assets/Scripts/ 目錄結構

這份文件概述了 Unity 專案中 `Assets/Scripts/` 目錄目前的組織方式。

```
Assets/Scripts/
├── Game/
│   ├── Models/
│   │   └── ConnectionModel.cs
│   ├── Presenters/
│   │   └── ConnectionPresenter.cs
│   ├── Singleton/
│   │   └── SystemManager.cs
│   ├── TestScript/
│   │   ├── AntnetEchoTest.cs
│   │   └── PositionUpdateTest.cs
│   └── Views/
│       ├── ConnectionView.cs
│       ├── IConnectionView.cs
│       └── TabManagerView.cs
└── Network/
    ├── Serializers/
    │   ├── ISerializer.cs
    │   ├── JSONSerializer.cs
    │   └── ProtobufSerializer.cs
    └── Transports/
        ├── IConnection.cs
        ├── TCPClientTransport.cs
        └── WebSocketClientTransport.cs
```

## 說明

* **Game/:** 包含遵循 MVP (Model-View-Presenter) 模式的核心遊戲邏輯組件，以及單例和測試腳本。
  * **Models/:** 資料和業務邏輯。
  * **Presenters/:** 協調視圖 (Views) 和模型 (Models) 之間的互動。
  * **Singleton/:** 包含用於全域存取和初始化的單例類別，例如 `SystemManager`。
  * **TestScript/:** 包含用於網路功能的各種測試腳本。
  * **Views/:** 與 UI 相關的組件和介面。
* **Network/:** 包含與網路通訊相關的組件。
  * **Serializers/:** 資料序列化（例如 Protobuf、JSON）的介面和實作。
  * **Transports/:** 網路傳輸層（例如 TCP、WebSocket）的介面和實作。
