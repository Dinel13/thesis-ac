import http from "k6/http";

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
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDM1Mzg2MjQsImV4cCI6MTY0MzYyNTAyNH0.BSn6wo2cq7ihwo1EKC4oHBnFDBU5qs8xBHctinIlx3Q",
    },
  };

  http.post(url, payload, params);
}
