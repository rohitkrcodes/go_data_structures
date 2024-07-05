main:
	go run main.go

test:
	go test -v -cover -short ./...

docker:
	docker build . -t go-list-ds:latest

dockerun:
	docker run go-list-ds:latest

.PHONY:
	main test docker dockerun