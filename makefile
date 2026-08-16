.PHONY: build build-web run test benchmark benchmark-lib clean

build:
	go build -o sluggo.bin .

build-web:
	mkdir -p web
	GOOS=js GOARCH=wasm go build -o web/game.wasm .
	find $$(go env GOROOT) -name "wasm_exec.js" -exec cp {} web/ \;

run:
	go run .

test:
	go test -race ./...

benchmark:
	make benchmark-lib

benchmark-lib:
	go test ./lib/ -bench=BenchmarkVector2 -benchmem

clean:
	rm -f sluggo.bin web/game.wasm web/wasm_exec.js
