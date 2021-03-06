import { sleep } from "k6";
import http from "k6/http";

export default function () {
  const url = `http://${__ENV.IP}:8080/krs/1`;
  const payload = JSON.stringify({
    id_mahasiswa: 1,
    mata_kuliahs: [
      {
        kode: "IF-100",
        nama: "Pemrograman script",
        sks: 3,
        dosen: "Bu Dina",
        semester: "Semester 6",
      },
      {
        kode: "IF-101",
        nama: "Pemrograman Berorientasi Objek",
        sks: 3,
        dosen: "Bu Nana",
        semester: "Semester 6",
      },
      {
        kode: "IF-102",
        nama: "Pemrograman Visual",
        sks: 3,
        dosen: "Pak Bimo",
        semester: "Semester 4",
      },
      {
        kode: "IF-103",
        nama: "Sistem Basis Data",
        sks: 4,
        dosen: "Pak Daro",
        semester: "Semester 8",
      },
      {
        kode: "IF-104",
        nama: "Perancangan Sistem Informasi",
        sks: 2,
        dosen: "Pak Bambang",
        semester: "Semester 8",
      }
    ],
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDA4NTMsImlzcyI6ImF1dGgiLCJ1c2VybmFtZSI6InNheWEifQ.rlhZFTpHCM8Wyk2CGD2t_Ozqks0xKa0Ndu4IOrfmb9k",
    },
  };

  http.put(url, payload, params);

  sleep(1)
}
