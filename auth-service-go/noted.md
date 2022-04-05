protoc proto/auth.proto --go_out=plugins=grpc:.

go tool pprof app_name cpu.prof

go tool pprof http://localhost:8081/debug/pprof/profile?seconds=10

png top web