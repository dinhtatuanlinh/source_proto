proto:
	rm -f pkg/operatorpb/*.go
	protoc --proto_path=proto/operator --go_out=pkg/operatorpb --go_opt=paths=source_relative --go-grpc_out=pkg/operatorpb \
	--go-grpc_opt=paths=source_relative proto/operator/*.proto
push:
	git push origin v1.0.7
.PHONY: proto push