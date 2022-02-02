import { sleep } from "k6";
import http from "k6/http";

export default function () {
  const url = `http://${__ENV.IP}:8082/verify/1`;

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  http.get(url, params);

  sleep(1)

}
