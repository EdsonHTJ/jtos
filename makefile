
OUT_NAME = jtos

# OS detection
ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
	OUT_NAME = jtos.exe
endif

all:
	go run ./cmd/jtos

build:
	go build  -o bin/jtos ./cmd/jtos

build-wasm: 
	GOARCH=wasm GOOS=js go build -o bin/jtos.wasm ./wasm

