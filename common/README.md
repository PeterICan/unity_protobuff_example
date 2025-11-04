# 📦 共用資料模型與封包格式

本目錄定義了 Server 和 Client 之間的通訊協定和資料結構。

---

## 📁 目錄結構

```
common/
├── proto/
│   ├── message.proto           # TCP 模式的 Protobuf 定義
│   └── json_api/
│       ├── common.api.proto    # WebSocket 模式的通用結構
│       ├── position.api.proto  # 位置相關訊息
│       └── gamer_info.api.proto # 玩家資訊相關訊息
├── generate_go_proto.bat       # Go 程式碼生成腳本
└── generate_csharp_proto.bat   # C# 程式碼生成腳本
```

**註**：生成的程式碼會分別放在：
- Go：`server/generated/` 目錄
- C#：`client/Assets/Scripts/Generated/` 目錄

---

## 🔄 兩種通訊模式的封包格式

### 模式 I：TCP + Protobuf

**封包結構**：
```
[Header: 12 bytes] + [Protobuf Data]
```

**Header 格式**：
```
Len   uint32 (4 bytes)  // 資料長度
Error uint16 (2 bytes)  // 錯誤碼
Cmd   uint8  (1 byte)   // 命令 (主要使用)
Act   uint8  (1 byte)   // 動作 (主要使用)
Index uint16 (2 bytes)  // 序號
Flags uint16 (2 bytes)  // 標記
```

**已定義的訊息**：
| Cmd | Act | 訊息類型 | 描述 |
|-----|-----|----------|------|
| 1 | 1 | `C2SPositionUpdate` | 客戶端位置更新 |
| 1 | 2 | `S2CPositionUpdate` | 伺服器位置回應 |

**註**：目前只有 `Cmd` 和 `Act` 欄位在使用中，其他欄位保留供未來擴展。

### 模式 II：WebSocket + JSON

**封包結構**：
```json
{
  "route": "position/update",
  "request_id": "uuid-string",
  "data": { ... }
}
```

**路由定義**：
| Route | 方向 | 訊息類型 | 描述 |
|-------|------|----------|------|
| `position/update` | C2S | `C2SPositionUpdate` | 位置更新請求 |
| `position/update` | S2C | `S2CPositionUpdate` | 位置更新回應 |
| `gamer_info/retrieve` | C2S | `C2SGamerInfoRetrieve` | 玩家資訊查詢 |
| `gamer_info/retrieve` | S2C | `S2CGamerInfoRetrieve` | 玩家資訊回應 |

---

## 📋 資料結構定義

### 位置相關

**PlayerPosition** (位置資料)：
```proto
message PlayerPosition {
  float x = 1;
  float y = 2;
  float z = 3;
  int64 timestamp = 4;
}
```

**C2SPositionUpdate** (位置更新請求)：
```proto
message C2SPositionUpdate {
  string route = 1;              // "position/update"
  string request_id = 2;         // 請求 ID
  PlayerPosition position = 3;   // 位置資料
}
```

### 玩家資訊相關

**GamerInfo** (玩家資訊)：
```proto
message GamerInfo {
  string nickname = 1;     // 暱稱
  int32 level = 2;         // 等級
  int32 money = 3;         // 金錢
  int32 gems = 4;          // 寶石
}
```

**C2SGamerInfoRetrieve** (玩家資訊查詢)：
```proto
message C2SGamerInfoRetrieve {
  string route = 1;        // "gamer_info/retrieve"
  string request_id = 2;   // 請求 ID
}
```

### 錯誤處理

**ErrorResponse** (錯誤回應)：
```proto
message ErrorResponse {
  string route = 1;         // 原始路由
  string request_id = 2;    // 原始請求 ID
  string error_code = 3;    // 錯誤代碼
  string error_message = 4; // 錯誤訊息
}
```

---

## 🔧 程式碼生成

### 生成 Go 程式碼
```bash
cd common
.\generate_go_proto.bat
```

### 生成 C# 程式碼
```bash
cd common
.\generate_csharp_proto.bat
```

---

## 📝 新增訊息類型流程

1. **定義 Protobuf**：在適當的 `.proto` 檔案中定義新訊息
2. **更新路由**：在 `json_api` 相關檔案中加入新的 `route` 字串
3. **生成程式碼**：執行生成腳本
4. **註冊處理器**：在 Server 端註冊新的訊息處理函式
5. **實作客戶端**：在 Client 端實作對應的發送/接收邏輯

這樣能確保兩種通訊模式都能正確處理新的訊息類型。