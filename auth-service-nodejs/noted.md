sudo npm i -g grpc-tools

protoc protos/auth.proto --js_out=import_style=commonjs,binary:.


P_REDIS=localhost npm start

node --prof

node --prof-process isolate-0x6375ac0-32948-v8.log > processed.txt

IP_REDIS=172.31.72.116 pm2 start index.js --update-env
sudo IP_REDIS=localhost  pm2 restart index.js --update-env