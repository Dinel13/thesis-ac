
protoc proto/krs.proto --go_out=plugins=grpc:.


IP_REDIS=localhost Q_PAY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo Q_KRS=https://sqs.ap-southeast-1.amazonaws.com/511415859101/krspay.fifo  go run main.go


export AWS_REGION=ap-southeast-1


[Unit]
Description=payment service
ConditionPathExists=/home/ubuntu/thesis-ac/payment-service
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitAS=infinity
LimitRSS=infinity
LimitCORE=infinity
LimitNOFILE=65536
WorkingDirectory=/home/ubuntu/thesis-ac/payment-service
ExecStart=/home/ubuntu/thesis-ac/payment-service/main
Environment=IP_REDIS=172.31.74.139
Environment=AWS_REGION=ap-southeast-1
Environment=AWS_SHARED_CREDENTIALS_FILE=/home/ubuntu/.aws/credentials
Environment=Q_PAY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo
Environment=Q_KRS=https://sqs.ap-southeast-1.amazonaws.com/511415859101/krspay.fifo
q

Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target