install:
	go mod vendor
	go mod tidy

start:
	go build -o main ./cmd/todo.go
	./main -p 3000