### run this service
``` bash
IP_KRS= IP_AUTH= IP_PAYMENT ./main
```

### run test jmx
./test -j=/home/din/Downloads/apache-jmeter-5.4.1/bin/jmeter -m="test" ./rest ./result/rest


ulimit -n 3072 


tambhkan di [service] di /etc/systemd/system/auth.service jika run pake systemd
LimitAS=infinity
LimitRSS=infinity
LimitCORE=infinity
LimitNOFILE=65536

IP_KRS=localhost IP_AUTH=localhost IP_PAYMENT=localhost go run main.go

IP_KRS=localhost IP_AUTH=localhost IP_PAYMENT=localhost Q_PAY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo Q_PROXY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/krspay.fifo go run main.go


IP_KRS=172.31.30.48 IP_AUTH=172.31.24.28 IP_PAYMENT=172.31.28.177 Q_PAY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo Q_PROXY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/krspay.fifo ./main 


[Unit]
Description=test service
ConditionPathExists=/home/ubuntu/thesis-ac/proxy
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitAS=infinity
LimitRSS=infinity
LimitCORE=infinity
LimitNOFILE=65536
WorkingDirectory=/home/ubuntu/thesis-ac/proxy
ExecStart=/home/ubuntu/thesis-ac/proxy/main
Environment=IP_KRS=172.31.30.48
Environment=IP_AUTH=https://lanjukang.com
Environment=IP_PAYMENT=172.31.28.177
Environment=Q_PAY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo
Environment=Q_PROXY=https://sqs.ap-southeast-1.amazonaws.com/511415859101/krspay.fifo

Restart=on-failure
RestartSec=10

[Install]
