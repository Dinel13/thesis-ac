const { startGrpc } = require("./grpc/app");
const { startRest } = require("./rest/app");

startGrpc();
startRest();
