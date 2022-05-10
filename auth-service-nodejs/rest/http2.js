const { readFileSync } = require("fs");
const http2 = require("http2");
const express = require("express");
const http2Express = require('http2-express-bridge')

const { login, verify, signup } = require("./controller");

const startRestServer = () => {
  // const app = express();
  const app = http2Express(express)

  app.use(express.json()); // for parsing application/json

  app.get("/", (req, res) => {
    res.send("Hello World!");
  });
  app.post("/signup", signup);
  app.post("/login", login);
  app.post("/verify", verify);

  const options = {
    key: readFileSync("./server.key"),
    cert: readFileSync("./server.crt"),
    allowHTTP1: true,
  };
  const server = http2.createSecureServer(options, app);
  server.listen(8081, () => console.log("REST server running at port 8081"));
};

module.exports = startRestServer;
