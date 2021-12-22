const jwt = require("jsonwebtoken");
const redis = require("../database/redis");

async function checkAccount(user) {
  let userExits = await redis.get(user.username);
  if (!user) {
    return { token: null };
  }
  userExits = JSON.parse(userExits);
  if (userExits.password !== user.password) {
    return { token: null };
  }
  token = jwt.sign(
    {
      username,
    },
    "secretKey@123",
    { expiresIn: "1d" }
  );
  return { token };
}

function verifyToken(token) {
  return jwt.verify(token, "secretKey@123");
}

function Login(call, callback) {
  console.log("Login");
  callback(null, checkAccount(call.request));
}

function Verify(call, callback) {
  console.log("VerifyToken");
  callback(null, verifyToken(call.request));
}

module.exports = {
   Login,
   Verify,
};