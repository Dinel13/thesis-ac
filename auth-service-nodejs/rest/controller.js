const jwt = require("jsonwebtoken");
const redis = require("../database/redis");

const signup = async (req, res) => {
  try {
    const { username, password } = req.body;
    const token = jwt.sign({ username }, "secretKey@123", { expiresIn: "1d" });
    await redis.set(
      username,
      JSON.stringify({ username, password }),
      "EX",
      60 * 60 * 24
    );
    res.status(201).json({ token });
  } catch (error) {
    console.log(error);
    res.status(500).json({
      message: error,
    });
  }
};

const login = async (req, res) => {
  try {
    console.log(req.body);
    const { username, password } = req.body;

    console.log(username, password);
    let user = await redis.get(username);
    console.log(user);
    if (!user) {
      return res.status(400).json({
        message: "Username tidak ditemukan",
      });
    }
    user = JSON.parse(user);
    if (password !== user.password) {
      return res.status(400).json({
        message: "Password salah",
      });
    }
    token = jwt.sign(
      {
        username,
      },
      "secretKey@123",
      { expiresIn: "1d" }
    );
    res.status(200).json({
      token,
    });
  } catch (error) {
    console.log(error);
    res.status(500).json({
      message: error,
    });
  }
};

const verify = async (req, res) => {
  try {
    const token = req.headers.authorization.split(" ")[1];
    console.log(token);
    if (!token) {
      return res.status(401).json({
        isAuth: false,
      });
    }

    const decoded = jwt.verify(token, "secretKey@123");
    console.log(decoded);
    if (!decoded.username) {
      console.log("gg");
      return res.status(401).json({
        isAuth: false,
      });
    }
    res.status(200).json({
      isAuth: true,
    });
  } catch (err) {
    return res.status(500).json({
      isAuth: false,
    });
  }
};

module.exports = {
  signup,
  login,
  verify,
};
