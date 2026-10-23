.PHONY: setup-debian-deps setup-audio-assets regenerate-audio build build-web serve-web run test test-ci benchmark benchmark-lib clean lint

setup-debian-deps:
	sudo apt-get update
	sudo apt-get install -y libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config xvfb binaryen

# extra deps for regenerating the embedded audio assets (fluidsynth, soundfont, ffmpeg, mido)
setup-audio-assets:
	sudo apt-get install -y fluidsynth fluid-soundfont-gm ffmpeg python3-mido

regenerate-audio:
	mkdir -p /tmp/sfx
	python3 tools/gen_sfx.py
	python3 tools/gen_music.py
	fluidsynth -ni -F /tmp/music.wav /usr/share/sounds/sf2/FluidR3_GM.sf2 /tmp/music.mid
	ffmpeg -y -loglevel error -i /tmp/music.wav -c:a libvorbis -q:a 3 assets/music.ogg
	for f in move eat over start; do ffmpeg -y -loglevel error -i /tmp/sfx/$$f.wav -c:a libvorbis -q:a 3 assets/sfx_$$f.ogg; done


lint:
	golangci-lint run ./...

build:
	CGO_ENABLED=0 go build -o sluggo.bin .

build-web:
	mkdir -p web
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o web/game.wasm .
	wasm-opt -Oz --enable-bulk-memory --enable-nontrapping-float-to-int --enable-sign-ext web/game.wasm -o web/game.wasm.opt
	mv web/game.wasm.opt web/game.wasm
	find $$(go env GOROOT) -name "wasm_exec.js" -exec cp {} web/ \;

serve-web:
	python3 web/devserver.py 9090

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
