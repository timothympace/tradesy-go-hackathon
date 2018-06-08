protoc: protoc-item protoc-user
protoc-item:
	protoc -I ./item-service/item --go_out=plugins=grpc:./item-service/item item.proto
protoc-user:
	protoc -I ./user-service/user --go_out=plugins=grpc:./user-service/user user.proto
