proto:
	rm -f pkg/operatorpb/*.go
	protoc --proto_path=proto/operator --go_out=pkg/operatorpb --go_opt=paths=source_relative --go-grpc_out=pkg/operatorpb \
	--go-grpc_opt=paths=source_relative proto/operator/*.proto
	rm -f pkg/contentpb/*.go
	protoc --proto_path=proto/content --go_out=pkg/contentpb --go_opt=paths=source_relative --go-grpc_out=pkg/contentpb \
	--go-grpc_opt=paths=source_relative proto/content/*.proto
	rm -f pkg/error_detailpb/*.go
	protoc --proto_path=proto/error_detail --go_out=pkg/error_detailpb --go_opt=paths=source_relative --go-grpc_out=pkg/error_detailpb \
    --go-grpc_opt=paths=source_relative proto/error_detail/*.proto
push:
	git add .
	git commit -m "fix"
	git tag v1.0.40
	git push origin v1.0.40
	git checkout master
	git merge v1.0.40
	git push origin master

.PHONY: proto push