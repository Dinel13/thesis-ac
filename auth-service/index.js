const grpc = require("grpc");
const express = require("express");
const { login, verify, signup } = require("./rest/controller");

function startGrpc() {
  const server = new grpc.Server();

  server.bind("127.0.0.1:50051", grpc.ServerCredentials.createInsecure());
  server.start();

  console.log("GRPC server running at port 50051");
}

function startRest() {
  const app = express();

  app.use(express.json()); // for parsing application/json

  app.post("/signup", signup);
  app.post("/login", login);
  app.post("/verify", verify);
  app.listen(3000, () => console.log("REST server running at port 3000"));
}

startGrpc();
startRest();
