### run this service
``` bash
IP_KRS= IP_AUTH= IP_PAYMENT ./main
```


./test -j=/home/din/Downloads/apache-jmeter-5.4.1/bin/jmeter -m="convert" ./rest ./report/rest

### run test jmx
./test -j=/home/din/Downloads/apache-jmeter-5.4.1/bin/jmeter -m="test" ./rest ./result/rest


ulimit -n 3072 