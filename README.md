# gRPC client side load balancing example using Go

This is a gRPC client-side load balancing example on top of [grpc-go](https://github.com/grpc/grpc-go).

The only load balancer bundled to grpc-go is `grpc.RoundRobin`. grpc RoundRoubin requires a `grpc.Resolver` which is intended to implement a DNS resolver or an other resourse resolution mechanism like Consul.

In this example I implemented a `grpc.Resolver` which only returns fixed servers initially passed. It is good for trying gRPC load balancing instantly.

## Prequirements

* Follow the instructions in https://grpc.io/docs/languages/go/quickstart/#prerequisites
* Install debug tools
  * Evans CLI
    * `go install github.com/ktr0731/evans@latest`
  * grpcurl
    * `curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.8.1/grpcurl_1.8.1_linux_x86_64.tar.gz | tar -xz`

## Run

```console
$ ./run.sh
```

# LICENSE
MIT @hakobe
