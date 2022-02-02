import { sleep } from "k6";
import http from "k6/http";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "krs.proto");


export default function () {
  const url = `http://${__ENV.IP}:8080/v2/krs`;
  const payload = JSON.stringify({
    id_mahasiswa: 1,
    mata_kuliahs: [
      {
        kode: "IF-100",
        nama: "Design Pattern",
        sks: 3,
        dosen: "Pak Ikbal",
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
      },
    ],
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDM2NzkzNzd9.CrsUKuqK9EEvwWeCVfexHGXi0sN-fauX5wmw1WnNtY8",
    },
  };

  http.post(url, payload, params);
  sleep(1);
}
