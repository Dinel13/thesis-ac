const express = require("express");
const { login, verify, signup } = require("./controller");

const  startRestServer = () => {
  const app = express();

  app.use(express.json()); // for parsing application/json

  app.post("/signup", signup);
  app.post("/login", login);
  app.post("/verify", verify);
  app.listen(8081, () => console.log("REST server running at port 8081"));
}

module.exports = startRestServer;
