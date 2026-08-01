.PHONY: build run test benchmark benchmark-lib check-gcc check-x11

check-gcc:
	@command -v gcc >/dev/null 2>&1 || { \
		echo "gcc not found."; \
		read -r -p "gcc is required to build. Install it now? (y/N) " ans; \
		case "$$ans" in \
			y|Y) \
				if command -v apt-get >/dev/null 2>&1; then \
					sudo apt-get update && sudo apt-get install -y gcc; \
				elif command -v dnf >/dev/null 2>&1; then \
					sudo dnf install -y gcc; \
				elif command -v pacman >/dev/null 2>&1; then \
					sudo pacman -S --noconfirm gcc; \
				elif command -v brew >/dev/null 2>&1; then \
					brew install gcc; \
				else \
					echo "Could not determine package manager. Please install gcc manually."; \
					exit 1; \
				fi ;; \
			*) echo "Aborting."; exit 1 ;; \
		esac; \
	}

check-x11:
	@if [ ! -f /usr/include/X11/Xlib.h ] || ! echo 'int main(){}' | gcc -x c - -lXxf86vm -o /dev/null 2>/dev/null; then \
		echo "X11 development headers/libraries not found."; \
		read -r -p "X11 dev headers and libs are required to build (Ebiten needs them). Install now? (y/N) " ans; \
		case "$$ans" in \
			y|Y) \
				if command -v apt-get >/dev/null 2>&1; then \
					sudo apt-get update && sudo apt-get install -y libx11-dev libxrandr-dev libxcursor-dev libxi-dev libxinerama-dev libxkbcommon-dev libgl1-mesa-dev libxxf86vm-dev; \
				elif command -v dnf >/dev/null 2>&1; then \
					sudo dnf install -y libX11-devel libXrandr-devel libXcursor-devel libXi-devel libXinerama-devel libxkbcommon-devel mesa-libGL-devel libXxf86vm-devel; \
				elif command -v pacman >/dev/null 2>&1; then \
					sudo pacman -S --noconfirm libx11 libxrandr libxcursor libxi libxinerama libxkbcommon mesa libxxf86vm; \
				elif command -v brew >/dev/null 2>&1; then \
					brew install --cask xquartz; \
				else \
					echo "Could not determine package manager. Please install the X11 dev headers manually."; \
					exit 1; \
				fi ;; \
			*) echo "Aborting."; exit 1 ;; \
		esac; \
	fi

build: check-gcc check-x11
	go build -o sluggo.bin .

run: check-gcc check-x11
	go run .

test: check-gcc check-x11
	go test -race ./...

benchmark:
	make benchmark-lib

benchmark-lib:
	go test ./lib/ -bench=BenchmarkVector2 -benchmem