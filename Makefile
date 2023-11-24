DATE=$(shell date +'%d-%m-%Y-%H:%M')

.PHONY: build

make:
	ENVIRONMENT=debug
	go run main.go

debug:
	ENVIRONMENT=debug
	go run main.go

serve:
	go run github.com/hajimehoshi/wasmserve@latest .

build:
	env GOOS=js GOARCH=wasm go build -o game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js .
build:
	env GOOS=js GOARCH=wasm go build -o game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js .
build_production:
	env GOOS=js GOARCH=wasm go build -o ./build/build-${DATE}/game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js ./build/build-${DATE}
	cp main.html ./build/build-${DATE}
	cp index.html ./build/build-${DATE}
