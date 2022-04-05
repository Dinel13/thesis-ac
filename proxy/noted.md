### run this service
``` bash
IP_KRS= IP_AUTH= IP_PAYMENT ./main
```


./test -j=/home/din/Downloads/apache-jmeter-5.4.1/bin/jmeter -m="convert" ./rest ./report/rest

### run test jmx
./test -j=/home/din/Downloads/apache-jmeter-5.4.1/bin/jmeter -m="test" ./rest ./result/rest


ulimit -n 3072 

IP_KRS=172.31.30.48 IP_AUTH=172.31.24.28 IP_PAYMENT=172.31.28.177 ./main 