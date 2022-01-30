import http from "k6/http";

export const options = {
  stages: [
    { duration: '1s', target: 1000 }, // simulate ramp-up of traffic from 1 to 100 users over 10 second.
    { duration: '1s', target: 0 }, // stay at 100 users for 1 minute
  ],
};

export default function () {
  const url = "http://127.0.0.1:8080/krs";
  const payload = JSON.stringify({
    id_mahasiswa: 2,
    mata_kuliahs: [
      {
        kode: "IF-141",
        nama: "Pemrograman script",
        sks: 3,
        dosen: "Dina",
        semester: "Semester 7",
      },
      {
        kode: "IF-101",
        nama: "Pemrograman Script",
        sks: 3,
        dosen: "Dina",
        semester: "Semester 7",
      },
    ],
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDA4NTMsImlzcyI6ImF1dGgiLCJ1c2VybmFtZSI6InNheWEifQ.rlhZFTpHCM8Wyk2CGD2t_Ozqks0xKa0Ndu4IOrfmb9k",
    },
  };

  http.post(url, payload, params);
}
