server:
	@go run .\main.go server
worker:
	@go run .\main.go worker
migrate:
	@go run .\main.go migrate
lint:
	@gofumpt -l -w .
	@golangci-lint run --config .golangci.yml
swagger:
	@swag init -g main.go
benchmark:
	@go build -o benchmark.exe ./cli && ./benchmark.exe
install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	go install mvdan.cc/gofumpt@latest