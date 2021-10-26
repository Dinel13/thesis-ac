sudo npm i -g grpc-tools

protoc protos/auth.proto --js_out=import_style=commonjs,binary:.