import http from 'k6/http';
import { sleep } from 'k6';

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
