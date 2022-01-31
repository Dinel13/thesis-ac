import http from "k6/http";

export const options = {
  stages: [
    { duration: '1s', target: 1 }, 
  ],
};

export default function () {
  const url = "http://127.0.0.1:8081/signup";
  const payload = JSON.stringify({
    username: "salahuddin",
    password: "Password123",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  http.post(url, payload, params);
}
