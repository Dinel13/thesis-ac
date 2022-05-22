# gRPC+Protobuf vs JSON+HTTP?

Inspired by [https://github.com/plutov/benchmark-grpc-protobuf-vs-http-json](https://github.com/plutov/benchmark-grpc-protobuf-vs-http-json



```
apt install protobuf-compiler-dev supaya langsung ada protoc-gen-go

```

## test
go test -bench=. -benchmem


## cpu

go test -bench=BenchmarkGRPCProtobuf -cpuprofile=grpcprotobuf.cpu
go test -bench=BenchmarkHTTPJSON -cpuprofile=httpjson.cpu

go tool pprof grpcprotobuf.cpu
go tool pprof httpjson.cpu