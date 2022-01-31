import http from "k6/http";

export const options = {
  stages: [
    { duration: '1s', target: 100 }, 
    { duration: '1s', target: 100 }, 
  ],
};

export default function () {
  const url = "http://127.0.0.1:8082/verify/1";

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  http.get(url, params);
}
