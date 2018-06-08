protoc-item:
	protoc -I ./item-service/item --go_out=plugins=grpc:./item-service/item item.proto
protoc-user:
	protoc -I ./user-service/user --go_out=plugins=grpc:./user-service/user user.proto
protoc-hellworld:
	protoc -I ./item-service/helloworld --go_out=plugins=grpc:./item-service/helloworld helloworld.proto