bindata.go: clean build app.js app.css
	cp html/* build/
	$$GOPATH/bin/go-bindata -pkg static -prefix build/ build/

clean:
	rm -f bindata.go
	rm -rf build

build:
	mkdir build

app.js:
	npm run build

app.css:
	npm run build-sass

prepare:
	go get -u github.com/jteeuwen/go-bindata/...
	npm install
	./node_modules/.bin/bower install

prepare-clean: clean
	rm -rf node_modules
	rm -rf bower_components
