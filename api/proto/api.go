package proto

//go:generate protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative --experimental_allow_proto3_optional --grpc-gateway_out . --grpc-gateway_opt paths=source_relative racing/racing.proto

//go:generate protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative --experimental_allow_proto3_optional --grpc-gateway_out . --grpc-gateway_opt paths=source_relative sports/sports.proto
