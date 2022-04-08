const Redis = require("ioredis");

// get ip for redis server from environment variable
const redisHost = process.env.IP_REDIS;

if (!redisHost) {
   throw new Error("IP_REDIS environment variable not set");
}

// create new redis client with given host
const redis = new Redis({
   host: redisHost,
   port: 6379,
   password: "",
   db: 0,
});

// cek connection
redis.on("connect", () => {
   console.log("Redis client connected to ", redisHost );
});

// handle error
redis.on("error", (err) => {
   console.log("Something went wrong " + err);
   throw new Error(err);
});


module.exports = redis;