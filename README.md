# Thesis-ac

My bachelor thesis in [Hasanuddin University](https://unhas.ac.id/v2/) as an Informatics Engineer Student.
<br>

## Table of contents

- [General info](#general-info)
- [Technologies](#technologies)
- [Run service](#run-service)
- [Run test](#run-test)
- [Result](#result)
  <br>

## General info

The main purpose of this thesis is to **Compare the performance of microservice use gRPC and REST APIs for academic management system**. So this project has 3 services:

1. **KRS Service** use golang.
2. **Auth Service** uses nodejs and also golang.
3. **Payment Service** use golang.

To store data in the database, I use **Redis**. Then use **k6** to test the performance of each service and use **Datadog** to monitor the performance. All the services are deployed in different **EC2 AWS** as well as the test agent then use private IP to communicate.
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
$ URL_AUTH=172.31.24.28 URL_PAYMENT=172.31.28.177 KRS-service/main
```

### 2. **Auth Service**

#### Use nodejs

```bash
$  cd auth-service-nodejs && npm install && npm start
```

#### Use golang

```bash
$ auth-service-go/main
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

Don't forget to replace _<YOUR_DATADOG_API_KEY>_ with your datadog api key
<br>

### 2. Run the test

All test files are in **test** directory. For example, if you want to test grpc-krs-create in 200vurtual user in 5 second use the command below.

```bash
 K6_STATSD_ENABLE_TAGS=true  k6 run --vus 200 --duration 5s -e IP=127.0.0.1 test/grpc/krs/create.js
```

<br>

## Result

Stil develop
