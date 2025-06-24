# ベースイメージ
FROM golang:1.23.0

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコピー
COPY . .

# 依存関係をインストール
RUN go mod tidy

# アプリケーションをビルド
RUN go build -o main .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

RUN chmod +x main

# アプリケーションを実行
CMD ["./main"]