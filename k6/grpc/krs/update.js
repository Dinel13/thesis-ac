import { sleep } from "k6";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "krs.proto");

export default () => {
  client.connect(`${__ENV.IP}:9090`, {
    plaintext: true,
  });

  const data = {
    token:
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDA4NTMsImlzcyI6ImF1dGgiLCJ1c2VybmFtZSI6InNheWEifQ.rlhZFTpHCM8Wyk2CGD2t_Ozqks0xKa0Ndu4IOrfmb9k",
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
  };
  
  client.invoke("proto.KrsService/Update", data);

  sleep(1)
};
