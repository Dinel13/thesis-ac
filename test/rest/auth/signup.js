import http from "k6/http";

export default function () {
  const url = `http://${__ENV.IP}:8081/signup`;
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
}
