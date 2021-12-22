const grpc = require("@grpc/grpc-js");
const PROTO_PATH = __dirname + "/../protos/auth.proto";
const protoLoader = require("@grpc/proto-loader");
const { Login, Verify } = require("./controller");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
// The protoDescriptor object has the full package hierarchy
const authservice = protoDescriptor.proto;

const server = new grpc.Server();

server.addService(authservice.AuthService.service, {
  Login: Login,
  Verify: Verify,
});

server.bindAsync(
  "0.0.0.0:50051",
  grpc.ServerCredentials.createInsecure(),
  () => {
    server.start();
  }
);
