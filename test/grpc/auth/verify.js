import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "auth.proto");

export default () => {
  client.connect("127.0.0.1:9091", {
    plaintext: true,
  });

  const data = {
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDA4NTMsImlzcyI6ImF1dGgiLCJ1c2VybmFtZSI6InNheWEifQ.rlhZFTpHCM8Wyk2CGD2t_Ozqks0xKa0Ndu4IOrfmb9k",
  };
  
  client.invoke("proto.AuthService/Verify", data);
  sleep(0.2);
};
