DATE=$(shell date +'%d-%m-%Y-%H:%M')

.PHONY: build

debug:
	cp -f config/env.debug config/env.go
	sudo env GOOS=js GOARCH=wasm go build -o ./build/debug/build-${DATE}/game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js ./build/debug/build-${DATE}/
dev:
	cp -f config/env.dev config/env.go
	env GOOS=js GOARCH=wasm go build -o ./build/dev/build-${DATE}/game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js ./build/dev/build-${DATE}/
prod:
	cp -f config/env.prod config/env.go
	env GOOS=js GOARCH=wasm go build -o ./build/prod/build-${DATE}/game.wasm github.com/seanburman/game
	cp /usr/local/go/misc/wasm/wasm_exec.js ./build/prod/build-${DATE}