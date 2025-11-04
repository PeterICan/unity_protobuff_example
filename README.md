# 🚀 應用層通訊抽象與協定實驗專案

本專案實作一套 **資料模型** 與 **通訊協定** 解耦的 Server/Client 架構，實驗兩種通訊模式：

| 模式 | 傳輸協定 | 資料格式 | 適用場景 |
| :--- | :--- | :--- | :--- |
| **模式 I** | **TCP** | **Protocol Buffers** | 即時數據交換、高頻率同步 |
| **模式 II** | **WebSocket** | **JSON** | 玩家資訊查詢、位置更新 |

## 技術棧

- **Server**: Go (使用 `antnet` 框架)
- **Client**: Unity (C#)

## 快速啟動

### 建置 Server

```bash
server\build_server.bat
```

### 建置 Unity Client  

```bash
client\build_project.bat
```

### Docker 部署 (可選)

```bash
server\build_docker_compose.bat
```

詳細說明請參考：

- 共用封包：`common/README.md`
- Server 端：`Server/README.md`
- Client 端：`Client/README.md`
