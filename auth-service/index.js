const grpc = require("grpc");
const express = require("express");
const { login, verify } = require("./rest/controller");

function startGrpc() {
  const server = new grpc.Server();

  server.bind("127.0.0.1:50051", grpc.ServerCredentials.createInsecure());
  server.start();

  console.log("Server running at port 50051");
}

function startRest() {
  const app = express();
  app.post("/login", login);
  app.post("/verify", verify);
  app.listen(3000, () => console.log("Example app listening on port 3000!"));
}

startGrpc();
startRest();
