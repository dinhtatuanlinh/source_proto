proto:
	rm -f pkg/operatorpb/*.go
	protoc --proto_path=proto/operator --go_out=pkg/operatorpb --go_opt=paths=source_relative --go-grpc_out=pkg/operatorpb \
	--go-grpc_opt=paths=source_relative proto/operator/*.proto
	rm -f pkg/contentpb/*.go
	protoc --proto_path=proto/content --go_out=pkg/contentpb --go_opt=paths=source_relative --go-grpc_out=pkg/contentpb \
	--go-grpc_opt=paths=source_relative proto/content/*.proto
push:
	git push origin v1.0.12
.PHONY: proto push