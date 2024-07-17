server:
	go run main.go

test:
	go test -v -cover -short ./...

runner:
	sh runner.sh

.PHONY:
	main test docker dockerun runner