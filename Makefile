DATE=$(shell date +'%d-%m-%Y-%H:%M')

.PHONY: build

make:
	ENVIRONMENT=debug
	go run main.go

debug:
	ENVIRONMENT=debug
	go run main.go

build_debug:
	cp -f config/env.debug config/env.go
	env GOOS=js GOARCH=wasm go build -o game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js .
build_dev:
	cp -f config/env.dev config/env.go
	env GOOS=js GOARCH=wasm go build -o game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js .
build_prod:
	cp -f config/env.prod config/env.go
	env GOOS=js GOARCH=wasm go build -o ./build/build-${DATE}/game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js ./build/build-${DATE}
	cp index.html ./build/build-${DATE}
