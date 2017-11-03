PROTO_PATH = proto/apis

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mcheckin/protobuf/empty.proto=$(PACKAGE)/apis/checkin/protobuf,

.PHONY: proto
proto: ## Generate code from protocol buffer
	@mkdir -p apis
	protoc -I ${PROTO_PATH} ${PROTO_PATH}/checkin/protobuf/empty.proto --go_out=$(PROTO_MAPPINGS)plugins=grpc:apis
	protoc -I ${PROTO_PATH} ${PROTO_PATH}/checkin/checkin/v1alpha/*.proto --go_out=$(PROTO_MAPPINGS)plugins=grpc:apis
	protoc -I ${PROTO_PATH} ${PROTO_PATH}/checkin/user/v1alpha/*.proto --go_out=$(PROTO_MAPPINGS)plugins=grpc:apis

envcheck::
	$(call executable_check,protoc,protoc)
	$(call executable_check,protoc-gen-go,protoc-gen-go)
