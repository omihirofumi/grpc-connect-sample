# grpc-connect-sample
### Usage
gRPCサーバーの起動
```terminal
go run ./cmd/server/main.go
または
make run
```

curlを使用したHttpリクエスト
```terminal
curl \
	--header "Content-Type: application/json" \
	--data '{"sentence": "Hello"}' \
	http://localhost:8080/chat.v1.ChatService/Say
```

grpcurlを使用したリクエスト
```terminal
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"sentence": "Hello"}' \
    localhost:8080 chat.v1.ChatService/Say
```
(grpcurlのインストール)
```terminal
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
