const startGrpc  = require("./grpc/app");
const startRest = require("./rest/http2");

startGrpc();
startRest();
