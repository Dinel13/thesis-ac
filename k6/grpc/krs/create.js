import grpc from "k6/net/grpc";
import { check, sleep } from "k6";

const client = new grpc.Client();
client.load(["./"], "krs.proto");

export const options = {
   stages: [
      { duration: '1s', target: 1000 }, // simulate ramp-up of traffic from 1 to 100 users over 10 second.
      { duration: '1s', target: 0 }, // stay at 100 users for 1 minute
    ],
 };

export default () => {
  client.connect("127.0.0.1:9090", {
    plaintext: true,
  });

  const data = {
    token:
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDA4NTMsImlzcyI6ImF1dGgiLCJ1c2VybmFtZSI6InNheWEifQ.rlhZFTpHCM8Wyk2CGD2t_Ozqks0xKa0Ndu4IOrfmb9k",
    id_mahasiswa: 2,
    mata_kuliahs: [
      {
        kode: "IF-141",
        nama: "Pemrograman Script",
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
  };
  const response = client.invoke("proto.KrsService/Create", data);

//   check(response, {
//     "status is OK": (r) => r && r.status === grpc.StatusOK,
//   });

//   console.log(JSON.stringify(response.message));

//   client.close();
//   sleep(1);
};
