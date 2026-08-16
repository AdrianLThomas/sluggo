.PHONY: setup-debian-deps build build-web run test test-ci benchmark benchmark-lib clean

setup-debian-deps:
	sudo apt-get update
	sudo apt-get install -y libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config xvfb

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

test-ci:
	xvfb-run -a make test

benchmark:
	make benchmark-lib

benchmark-lib:
	go test ./lib/ -bench=BenchmarkVector2 -benchmem

clean:
	rm -f sluggo.bin web/game.wasm web/wasm_exec.js
