import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "auth.proto");

export const options = {
  stages: [
    { duration: "1s", target: 100 },
    { duration: "1s", target: 100 },
  ],
};

export default () => {
  client.connect("127.0.0.1:9091", {
    plaintext: true,
  });

  const data = {
    username: "salahuddin",
    password: "Password123",
  };
  
  client.invoke("proto.AuthService/Login", data);
};
