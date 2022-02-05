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
  };

  client.invoke("proto.KrsService/Read", data);

  sleep(1)
};
