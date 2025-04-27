proto:
	rm -f pkg/operatorpb/*.go
	protoc --proto_path=proto/operator --go_out=pkg/operatorpb --go_opt=paths=source_relative --go-grpc_out=pkg/operatorpb \
	--go-grpc_opt=paths=source_relative proto/operator/*.proto
push:
	git tag v1.0.2
	git push origin v1.0.2
.PHONY: proto push