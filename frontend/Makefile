# shell is set so new path is picked up
SHELL := /bin/bash
export PATH := ./node_modules/.bin:$(GOPATH)/bin:$(PATH)

# directories
SRC_DIR=./src
TEST_DIR=./test
BUILD_DIR=./build

include scripts/prepare
include scripts/build
include scripts/test

# clean ALL the things
clean: build-clean prepare-clean test-clean