
const express = require("express");
const { login, verify, signup } = require("./controller");

export function startRest() {
   const app = express();
 
   app.use(express.json()); // for parsing application/json
 
   app.post("/signup", signup);
   app.post("/login", login);
   app.post("/verify", verify);
   app.listen(3000, () => console.log("REST server running at port 3000"));
 }