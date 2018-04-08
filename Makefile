all: build
build-server:
	go build -o server/server github.com/hakobe/grpc-try-load-balancing/server
build-client:
	go build -o client/client github.com/hakobe/grpc-try-load-balancing/client
gen-pb:
	mkdir -p trylb && protoc ./trylb.proto --go_out=plugins=grpc:trylb
build: gen-pb build-server build-client

.PHONY: \
	build-server \
	build-client \
	gen-pb \
	build \
	all
