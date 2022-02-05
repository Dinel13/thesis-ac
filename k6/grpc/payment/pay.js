import { sleep } from "k6";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(["./"], "auth.proto");

export default () => {
  client.connect(`${__ENV.IP}:9092`, {
    plaintext: true,
  });

  const data = {
    id_mahasiswa: 1,
    jumlah: 75000,
    metode: "transfer BNI",
  };
  
  client.invoke("proto.PaymentService/Create", data);

  sleep(1)
};
