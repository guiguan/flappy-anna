GO := $(GOM)
DIST := dist

all: build

.PHONY: build http convert

build:
	GOOS=js GOARCH=wasm $(GO) build -pkgdir=vendor -o $(DIST)/flappy.wasm github.com/guiguan/flappy-anna
http:
	$(GO) build github.com/guiguan/flappy/cmd/http
convert:
	$(GO) build github.com/guiguan/flappy/cmd/convert
