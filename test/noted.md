import http from 'k6/http';
import { sleep } from 'k6';

--no-connection-reuse 

When running the test, you can use iftop in the terminal to view in real-time the amount of network traffic generated. If the traffic is constant at 1Gbit/s, your test is probably limited by the network card. Consider upgrading to a different EC2 instance.

const res = http.get('https://test.k6.io');
const checkRes = check(res, {
  'Homepage body size is 11026 bytes': (r) => r.body && r.body.length === 11026,
});
COPY
Code like this runs fine when the system under test (SUT) is not overloaded and returns proper responses. When the system starts to fail, the above check won't work as expected.

The issue here is that the check assumes that there's always a body in a response. The r.body may not exist if server is failing. In such case, the check itself won't work as expected and error similar to the one below will be returned:

// default option tidak pelru ketik di terminal
export const options = {
   vus: 10,
   duration: '30s',
 };

export default function () {
  http.get('https://test.k6.io');
  sleep(1);
}



sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
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
