SHELL = /bin/sh
.PHONY: sum sub

sum:
	cd sum
	protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. sum.proto
	sed -i 's/__/sum/g' ./sum.pb.go
	sed -i 's/__/sum/g' ./sum_grpc.pb.go
	go mod init github.com/TJFG1998/microservices/sum
	go mod tidy

sub:
	cd sub
	protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. sub.proto
	sed -i 's/__/sub/g' ./sub.pb.go
	sed -i 's/__/sub/g' ./sub_grpc.pb.go
	go mod init github.com/TJFG1998/microservices/sub
	go mod tidy

mul:
	cd mul
	protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. mul.proto
	sed -i 's/__/mul/g' ./mul.pb.go
	sed -i 's/__/mul/g' ./mul_grpc.pb.go
	go mod init github.com/TJFG1998/microservices/mul
	go mod tidy

div:
	cd div
	protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. div.proto
	sed -i 's/__/div/g' ./div.pb.go
	sed -i 's/__/div/g' ./div_grpc.pb.go
	go mod init github.com/TJFG1998/microservices/div
	go mod tidy