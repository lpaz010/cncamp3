export tag=v1.0
root:
	export ROOT=github.com/cncamp3

build:
	echo "building httpServer binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpServer container"
	docker build -t cncamp3/module2/httpServer:${tag} .

push: release
	echo "pushing cncamp3/module/httpServer"
	docker push cncamp3/module2/httpServer:${tag}