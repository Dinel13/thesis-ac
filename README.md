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

To store data in the database, I use **Redis**. Then use **k6** to test the performance of each service and use **Datadog** to monitor the performance. Additionaly i also try use **Jmeter** to load test. All the services are deployed in different **EC2 AWS** as well as the test agent then use private IP to communicate.
<br>

**System design**
![sistem design](https://github.com/dinel13/thesis-ac/blob/main/design.jpg?raw=true)

## Technologies

- **Golang 1.17.1**
- **Nodejs 4.17.6**
- **Redis 5.0.7**
- **EC2 AWS**
- **K6**
- **Datadog 0.36.0**
- **Jmeter 5.4.1**
- **Linux ubuntu 20.4**
  <br>

## Run service

### 1. **KRS Service**

```bash
$ URL_AUTH=127.0.0.1 URL_PAYMENT=127.0.0.1 KRS-service/main
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

### 1. Use K6

#### Run DataDog agent

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

#### Run test use K6

All test files for K6 are in **k6** directory. For example, if you want to test grpc-krs-create in 100 vurtual user in 30 second use the command below.

```bash
 K6_STATSD_ENABLE_TAGS=true k6 run --vus 100 --durations 30s --out statsd --tag test_run_id=1 -e IP=172.31.30.48 k6/grpc/krs/create.js

```

<br>

### 2. Use Jmeter

#### Run gateway service

Gateway service use to transform from rest to grpc and carry on rest to rest. So we can use jmeter to test the performance both for grpc and rest.

```bash
$ test/main
```

#### Run test use JMeter

All jmx test file for Jmeter are in **jmx** directory. For example, if you want to test grpc-krs-create in 100 thread second use the command below.

```bash
jmeter -n -t jmx/100/grpc-krs-create.jmx -l testresults.jtl

```

### 3. Use Locust

Still developing...

<br>

## Result

### Result for K6 with 100 vurtual user in 30 second

**Grpc**
![grpc](https://github.com/dinel13/thesis-ac/blob/main/100.png?raw=true)
<br>

**Rest**
![rest](https://github.com/dinel13/thesis-ac/blob/main/100r.png?raw=true)
<br>

**Datadog**
![datadog](https://github.com/dinel13/thesis-ac/blob/main/datadog.png?raw=true)
