all: clean-build build app.js app.css
	cp html/index.html build/
	$$GOPATH/bin/go-bindata -pkg static -prefix build/ build/

clean-build:
	rm bindata.go
	rm -rf build
	mkdir build

app.js:
	npm run build

app.css:
	npm run build-sass

prepare:
	go get -u github.com/jteeuwen/go-bindata/...
	npm install
	./node_modules/.bin/bower install

clean: clean-build
	rm -rf node_modules
	rm -rf bower_components
