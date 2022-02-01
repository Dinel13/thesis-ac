import { sleep } from "k6";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "auth.proto");

export default () => {
  client.connect(`${__ENV.IP}:9091`, {
    plaintext: true,
  });

  const data = {
    username: "salahuddin",
    password: "Password123",
  };

  const res = client.invoke("proto.AuthService/Login", data);
  // console.log(JSON.stringify(res.message));
  sleep(0.2);
};
