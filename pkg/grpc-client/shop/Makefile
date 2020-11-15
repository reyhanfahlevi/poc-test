.PHONY: pb
pb:
	@protoc --proto_path=${GOPATH}/src --proto_path=${GOPATH}/src/github.com/golang/protobuf -I proto/ proto/*.proto --gofast_out=plugins=grpc:proto --govalidators_out=plugins=grpc:proto --doc_out=. --doc_opt=markdown,doc.md proto/*.proto
