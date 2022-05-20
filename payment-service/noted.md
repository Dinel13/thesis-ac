
protoc proto/krs.proto --go_out=plugins=grpc:.


IP_REDIS=localhost Q_PAY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo Q_KRS=https://sqs.ap-southeast-1.amazonaws.com/511415859101/krspay.fifo  go run main.go