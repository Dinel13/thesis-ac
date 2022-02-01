import { sleep } from "k6";
import grpc from "k6/net/grpc";

const client = new grpc.Client();

export default () => {
  client.connect(`${__ENV.IP}:9091`, {
    plaintext: true,
    reflect: true,
  });

  const res = client.invoke("proto.AuthService/Login", {
    username: "salahuddin",
    password: "Password123",
  });
  console.log(JSON.stringify(res.message));
  sleep(0.2);
};
