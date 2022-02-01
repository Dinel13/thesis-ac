import { sleep } from "k6";
import http from "k6/http";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "auth.proto");

export default function () {
  const url = `http://${__ENV.IP}:8081/login`;
  const payload = JSON.stringify({
    username: "salahuddin",
    password: "Password123",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.post(url, payload, params);
  console.log(res.body);
  sleep(0.2);
}
