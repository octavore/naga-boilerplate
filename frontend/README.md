# Naga Frontend

This module serves static assets. The static assets are built using Typescript and SASS into the `./build` folder, and then `go-bindata` is used to generate a Go file with all assets embedded into it. The Naga module in this folder is responsible for mapping HTTP requests to the embedded files.

The Makefile contains rules for building the `bindata.go` file, and the sub-commands for compiling the Typescript and SASS code are in `package.json`.

`package.json`, `bower.json`, `tsconfig.json`, and `tslint.json` contain various settings for asset compilation, mostly pertaining to Typescript.

## Usage

```
make prepare # installs go-bindata, and node and bower deps.
make sources # builds all sources into ./build
make test    # builds tests into ./test and runs them
```
