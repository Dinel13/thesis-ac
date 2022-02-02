import { sleep } from "k6";
import http from "k6/http";

export default function () {
  const url = `http://${__ENV.IP}:8082/pay`;
  const payload = JSON.stringify({
    id_mahasiswa: 1,
    jumlah: 75000,
    metode: "transfer BNI",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
    Authorization:
      "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDA4NTMsImlzcyI6ImF1dGgiLCJ1c2VybmFtZSI6InNheWEifQ.rlhZFTpHCM8Wyk2CGD2t_Ozqks0xKa0Ndu4IOrfmb9k",
  };

  http.post(url, payload, params);

  sleep(1)
}
