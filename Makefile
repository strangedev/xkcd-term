CC := go build

ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

PHONY: xkcd.sha256

clean:
	rm build/*

xkcd: cmd/feed/main.go src/feed.go
	$(CC) -o build/xkcd cmd/feed/main.go

xkcd.sha256: xkcd
	cd build && sha256sum xkcd > xkcd.sha256

install: build/xkcd
	install build/xkcd $(PREFIX)/bin
