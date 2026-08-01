.PHONY: build run test benchmark benchmark-lib

build:
	go build -o sluggo.bin .

run:
	go run .

test:
	go test -race ./...

benchmark:
	make benchmark-lib

benchmark-lib:
	go test ./lib/ -bench=BenchmarkVector2 -benchmem