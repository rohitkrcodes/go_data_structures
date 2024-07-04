main:
	go run main.go

test:
	go test -v -cover -short ./...

.PHONY:
	main test