sudo npm i -g grpc-tools

protoc protos/auth.proto --js_out=import_style=commonjs,binary:.

node --prof

node --prof-process isolate-0x6375ac0-32948-v8.log > processed.txt