# thesis-ac

My becholor thesis in [Hasanuddin University](https://unhas.ac.id/v2/) as informatics engineer.
<br>

## Table of contents

- [General info](#general-info)
- [Technologies](#technologies)
- [Run service](#run-service)
- [Run test](#run-test)
- [Result](#result)
<br>

## General info

The main purpose of this thesis is to **Comparing the performance of microservice use gRPC and REST APIs for acedemic management system**. So this project have 3 services:

1. **KRS Service** use golang.
2. **Auth Service** use nodejs and also golang.
3. **Payment Service** use golang.

To store data in database, I use **Redis**. Then use **k6** to test the performance of each service and use **datadog** to monitor the performance. All the services are deploy in diffent **ec2 aws** as well as the test agent then use private ip to comunicate.
<br>

## Technologies

- **Golang 1.17.1**
- **Nodejs 4.17.6**
- **Redis 5.0.7**
- **EC2 AWS**
- **K6**
- **Datadog 0.36.0**
- **Linux ubuntu 20.4**
<br>

## Run service

### 1. **KRS Service**

```bash
$ KRS-service/main
```
<br>

### 2. **Auth Service** use nodejs.

#### Use nodejs
```bash
$  cd auth-service && npm install && npm start
```

#### Use golang
```bash
$ auth-service/main
```
<br>

### 3. **Payment Service** use golang.

```bash
$ payment-service/main
```
<br>

## Run test

### 1. Run datadog agent

```bash
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
```
dont forget to replace *<YOUR_DATADOG_API_KEY>* with your datadog api key

<br>
### 2. Run the test
   
   All test file are in k6 directory. for example if you want to test grpc-krs-create in 100 request use the command below.
```bash
 K6_STATSD_ENABLE_TAGS=true  k6 run --out statsd k6/100/grpc/krs/create.js
```
<br>

## Result