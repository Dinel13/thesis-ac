# todo



# protoc

protoc --proto_path=. --python_out=. auth.proto
protoc --proto_path=proto --java_out=. proto/krs.proto
protoc proto/krs.proto --go_out=plugins=grpc:.

# scp
sudo scp -i /home/din/Documents/skripsi/skripsi.pem ubuntu@18.141.203.234:/home/ubuntu/thesis-ac/grpc-auth-1000.jtl ./

sudo scp -i Documents/skripsi/skripsi.pem ubuntu@3.1.203.66:/home/ubuntu/thesis-ac/auth-go/grpc/100/login1.jtl ./

sudo scp -i Documents/skripsi/skripsi.pem -r  ubuntu@3.1.203.66:/home/ubuntu/thesis-ac/auth-go ./

# ghz

../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.GetKrs   -d '{"id":1}' -n 10000 -c 1000   0.0.0.0:8081

ghz/ghz --insecure   --proto ./test/grpc/auth/auth.proto  --call  proto.AuthService.Login   -d '{"username": "salahuddin", "password":"Password123"}' -n 10000 -c 1000   0.0.0.0:9091

create krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Create   -d '{"token":       "ffdafa","id_mahasiswa":1,"mata_kuliahs" : [{"kode":     "IF-141","nama":"Pemrograman script", "sks":      3, "dosen":    "Dina","semester": "Semester 7"},{"kode":     "IF-101","nama":     "Pemrograman Script","sks":      3,"dosen":    "Dina","semester": "Semester 7"}]}' -n 1 -c 1  0.0.0.0:9090

create payment
../ghz/ghz --insecure   --proto ./proto/payment.proto   --call  proto.PaymentService.Create   -d '{"id_mahasiswa":1,"jumlah" : 2345, "metode" : "bri"}' -n 1 -c 1   0.0.0.0:9092

update krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Update   -d '{"token":       "ffdafa","id_mahasiswa":1,"mata_kuliahs" : [{"kode":     "IF-141","nama":"Pemrograman script", "sks":      3, "dosen":    "Dina","semester": "Semester 7"},{"kode":     "IF-101","nama":     "Pemrograman Script","sks":      3,"dosen":    "Dina","semester": "Semester 7"}]}' -n 1 -c 1   0.0.0.0:9090

read krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Read   -d '{"token":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDI2ODE5OTJ9.hMVVyeZyWg5Jq7gLiM3uKyV5jA_GbwZastQ2CWJzqek","id_mahasiswa":1}' -n 1 -c 1   0.0.0.0:9090


delete krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Delete   -d '{"token":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDI2ODE5OTJ9.hMVVyeZyWg5Jq7gLiM3uKyV5jA_GbwZastQ2CWJzqek","id_mahasiswa":1}' -n 1 -c 1   0.0.0.0:9090


# redis
reddis https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-20-04-id

sudo apt update
sudo apt install redis-server
sudo nano /etc/redis/redis.conf
supervised systemd
sudo systemctl restart redis.service
sudo systemctl status redis

# docker
sudo apt install docker.io

# jmeter
install jmeter
- pastika nsudah instal jvm
sudo apt install openjdk-11-jdk


wget https://downloads.apache.org//jmeter/binaries/apache-jmeter-5.4.1.tgz
tar -xvzf apache-jmeter-5.4.1.tgz
echo 'export PATH="$PATH:/home/ubuntu/apache-jmeter-5.4.1/bin"' >> ~/.bashrc
source ~/.bashrc

jmeter.sh -n -t jmx/100/grpc-krs-create.jmx -l testresults.jtl
./bin/jmeter  -g grpc.jtl -o reportt


# K6
echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
sudo apt-get update
sudo apt-get install k6

k6 run script.js
k6 run --vus 10 --duration 30s script.js
k6 run --vus 10 --duration 1s -e IP=127.0.0.1 test/rest/auth/login.js 

DOCKER_CONTENT_TRUST=1 \
sudo docker run -d \
    --name datadog \
    -v /var/run/docker.sock:/var/run/docker.sock:ro \
    -v /proc/:/host/proc/:ro \
    -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
    -e DD_SITE="datadoghq.com" \
    -e DD_API_KEY=<YOUR_DATADOG_API_KEY> \
    -e DD_DOGSTATSD_NON_LOCAL_TRAFFIC=1 \
    -p 8125:8125/udp \
    datadog/agent:latest

sudo docker stop newrelic-statsd 
sudo docker rm newrelic-statsd 
sudo docker logs newrelic-statsd 

k6 run --vus 1 --duration 1s -e IP=172.31.24.28 test/rest/auth/signup.js


K6_STATSD_ENABLE_TAGS=true  k6 run --out statsd --tag test_run_id=2 k6/rest/krs/create.js 


tekan g di dashbord untuk lihat metric cpu

k6 run --vus 10  -e IP=127.0.0.1 test/rest/krs/create.js 


K6_STATSD_ENABLE_TAGS=true k6 run --vus 100 --iterations 100 --out statsd --tag test_run_id=1 -e IP=172.31.30.48 test/grpc/krs/create.js

# locust
// generate 2 file py
virtualenv -p python3 env
source env/bin/activate
pip install grpcio grpcio-tools
python -m grpc_tools.protoc --proto_path=. ./auth.proto --python_out=. --grpc_python_out=.