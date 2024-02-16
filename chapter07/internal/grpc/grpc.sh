#!/usr/bin/env bash
# protoc --go_out=plugins=grpc:. greeter/greeter.proto  

# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greeter/greeter.proto

 protoc --go_out=. --go-grpc_out=. greeter/greeter.proto