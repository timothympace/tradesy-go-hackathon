protoc: protoc-item protoc-user protoc-search
protoc-item:
	protoc -I ./item-service/item --go_out=plugins=grpc:./item-service/item item.proto
protoc-user:
	protoc -I ./user-service/user --go_out=plugins=grpc:./user-service/user user.proto
protoc-search:
	protoc -I ./search-service/search --go_out=plugins=grpc:./search-service/search search.proto
kill:
	pkill -f go-build
	pkill -f go 