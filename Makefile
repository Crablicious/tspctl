all: build

build:
	go build

.PHONY: tags
tags:
	ctags --languages=go -e -R .
