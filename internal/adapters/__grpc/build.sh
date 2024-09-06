protoc -I=. --go_out=. --go-grpc_out=. api.proto

cp ./grpcapi/*.go ../../../../terminal_go/internal/hub/grpcapi