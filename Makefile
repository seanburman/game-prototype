DATE=$(shell date +'%d-%m-%Y-%H:%M')

.PHONY: build

debug:
	cp -f config/env.debug config/env.go
	sudo env GOOS=js GOARCH=wasm go build -o ./build/debug/build-${DATE}/game.wasm github.com/seanburman/game
dev:
	cp -f config/env.dev config/env.go
	sudo env GOOS=js GOARCH=wasm go build -o ./build/dev/build-${DATE}/game.wasm github.com/seanburman/game
prod:
	cp -f config/env.prod config/env.go
	sudo env GOOS=js GOARCH=wasm go build -o ./build/prod/build-${DATE}/game.wasm github.com/seanburman/game

server:
	sudo env GOOS=js GOARCH=wasm go build -o ../game-ws-server/static/game.wasm github.com/seanburman/game