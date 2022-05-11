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



IP_KRS=172.31.30.48 IP_AUTH=172.31.24.28 IP_PAYMENT=172.31.28.177 ./main 