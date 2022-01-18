const jwt = require("jsonwebtoken");
const redis = require("../database/redis");

function checkAccount(user) {}

async function Login(_, callback) {
  let userExits = await redis.get(_.request.username);
  userExits = JSON.parse(userExits);

  if (userExits.password !== _.request.password) {
    return { token: null };
  }
  token = jwt.sign(
    {
      username,
    },
    "secretKey@123",
    { expiresIn: "1d" }
  );
  callback(null, { token });
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
  callback(null, respon);
}

module.exports = {
  Login,
  Verify,
};
