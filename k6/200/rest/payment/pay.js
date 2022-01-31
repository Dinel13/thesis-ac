import http from "k6/http";

export const options = {
  stages: [
    { duration: '1s', target: 100 }, 
    { duration: '1s', target: 100 }, 
  ],
};

export default function () {
  const url = "http://127.0.0.1:8082/pay";
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
}
