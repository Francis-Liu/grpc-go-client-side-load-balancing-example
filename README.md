# gRPC client side load balancing example using Go

This is a gRPC client-side load balancing example on top of [grpc-go](https://github.com/grpc/grpc-go).

The only load balancer bundled to grpc-go is `grpc.RoundRobin`. grpc RoundRoubin requires a `grpc.Resolver` which is intended to implement a DNS resolver or an other resourse resolution mechanism like Consul.

In this example I implemented a `grpc.Resolver` which only returns fixed servers initially passed. It is good for trying gRPC load balancing instantly.

## Prequirements

- Follow the instructions in https://grpc.io/docs/languages/go/quickstart/#prerequisites

## Run

```console
# Run 4 servers
$ ./server/server -hostport 0.0.0.0:5000 &
$ ./server/server -hostport 0.0.0.0:5001 &
$ ./server/server -hostport 0.0.0.0:5002 &
$ ./server/server -hostport 0.0.0.0:5003 &

$ sleep 3 # Wait for server to start up

# Do gRPC method calls 10000 times
$ time ./client/client -n 10000 \
    -server localhost:5000 \
    -server localhost:5001 \
    -server localhost:5002 \
    -server localhost:5003 \
 
# or 
$ ./run.sh
```

# LICENSE
MIT @hakobe
