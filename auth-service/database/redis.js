const Redis = require("ioredis");
const redis = new Redis(); // uses defaults unless given configuration object

module.exports = redis;