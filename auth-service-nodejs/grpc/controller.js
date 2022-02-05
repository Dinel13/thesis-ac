const jwt = require("jsonwebtoken");
const redis = require("../database/redis");

function checkAccount(user) {}

async function Login(_, callback) {
  let token;
  try {
    let userExits = await redis.get(_.request.username);
    if (!userExits) {
      throw new Error("User not exits");
    }
    userExits = JSON.parse(userExits);

    if (userExits.password != _.request.password) {
      throw new Error("Password not match");
    }
    token = jwt.sign(
      {
        username: userExits.username,
      },
      "secretKey@123",
      { expiresIn: "1d" }
    );
  } catch (error) {
    console.log(error);
    return callback(null, { token: "" });
  }
  return callback(null, { token });
}

async function Verify(_, callback) {
  let respon;
  try {
    const user = await jwt.verify(_.request.token, "secretKey@123");
    if (user) {
      respon = {
        is_auth: true,
      };
    } else {
      respon = {
        is_auth: false,
      };
    }
  } catch (error) {
    respon = {
      is_auth: false,
    };
  }
  return callback(null, respon);
}

module.exports = {
  Login,
  Verify,
};
