SHELL := /bin/bash
export PATH := $(GOPATH)/bin:$(PATH)

default: clean emails bindata

emails: build $(subst src,build,$(subst .mjml,.html,$(wildcard src/*.mjml)))

build:
	mkdir build

build/%.html: src/%.mjml
	./node_modules/.bin/mjml -r $< -o $@

clean:
	rm -rf build/*.html
	rm bindata.go

bindata:
	go-bindata -pkg emails build
