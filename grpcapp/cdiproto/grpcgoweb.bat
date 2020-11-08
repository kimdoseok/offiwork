protoc -I=. cdidemo.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
protoc -I . cdidemo.proto --go_out=plugins=grpc:.