const jwt = require("jsonwebtoken");

const login = async (req, res) => {
  const { username, password } = req.body;
  const user = await User.findOne({ username });
  if (!user) {
    return res.status(400).send("Invalid username or password");
  }
  const isMatch = await user.comparePassword(password);
  if (!isMatch) {
    return res.status(400).send("Invalid username or password");
  }
  const token = jwt.sign({ _id: user._id }, process.env.SECRET);
  res.header("auth-token", token).send(token);
};

const verify = async (req, res) => {
  const token = req.headers["Authorization"];

  if (!token) return res.status(401).send("Access denied. No token provided.");
  try {
    const decoded = jwt.verify(token, process.env.SECRET);
    req.user = decoded;
    next();
  } catch (ex) {
    res.status(400).send("Invalid token.");
  }
};

module.exports = {
  login,
  verify,
};
