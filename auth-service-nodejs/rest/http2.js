const http2Express = require("http2-express-bridge");

const { login, verify, signup } = require("./controller");

const startRestServer = () => {
  // const app = express();
  const app = http2Express(express);

  app.use(express.json()); // for parsing application/json

  app.use(express.static(__dirname, { dotfiles: "allow" }));

  app.get("/", (req, res) => {
    res.send("Hello World!");
  });
  app.post("/signup", signup);
  app.post("/login", login);
  app.post("/verify", verify);

  const options = {
    key: readFileSync(
      "/etc/letsencrypt/live/lanjukang.com/privkey.pem",
      "utf8"
    ),
    cert: readFileSync("/etc/letsencrypt/live/lanjukang.com/cert.pem", "utf8"),
    ca: readFileSync(
      "/etc/letsencrypt/live/lanjukang.com/fullchain.pem",
      "utf8"
    ),
    allowHTTP1: true,
  };
  
  const server = http2.createSecureServer(options, app);
  server.listen(443, () => console.log("REST server running at port 443"));
};

module.exports = startRestServer;
