# 使用 Go 基礎映像來建置
FROM golang:1.24-alpine AS builder

# 設定容器內的工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 以下載依賴
COPY go.mod ./go.mod
COPY go.sum ./go.sumgo.mod
COPY go.sum ./go.sum
RUN go mod download

# 複製應用程式的其餘原始碼
COPY . .

# 建置 Go 應用程式
# CGO_ENABLED=0 用於靜態編譯，使二進位檔自包含
# -o 指定輸出檔名
# ./... 建置當前目錄及其子目錄中的所有套件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server_executable .

# 使用最小的基礎映像作為最終階段
FROM alpine:latest

# 設定工作目錄
WORKDIR /root/

# 從建置階段複製已建置的執行檔
COPY --from=builder /app/server_executable .

# 暴露伺服器監聽的埠
EXPOSE 6666
EXPOSE 7777

# 運行執行檔的命令
CMD ["./server_executable", "--serverType", "ws"]