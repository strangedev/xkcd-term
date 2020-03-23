CC := go build

ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

PHONY: xkcd

clean:
	rm build/*

xkcd: cmd/feed/main.go src/feed.go
	$(CC) -o build/xkcd cmd/feed/main.go

install: build/xkcd
	install build/xkcd $(PREFIX)/bin
