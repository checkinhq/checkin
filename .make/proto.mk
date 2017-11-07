PROTO_PATH = proto/apis
PROTO_INCLUDE = -I ${PROTO_PATH} -I proto/protobuf/src -I vendor

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mcheckin/protobuf/empty.proto=$(PACKAGE)/apis/checkin/protobuf,

WEB_PROTO_MAPPINGS :=
WEB_PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mcheckin/protobuf/empty.proto=$(PACKAGE)/web/apis/checkin/protobuf,

.PHONY: proto
proto: ## Generate code from protocol buffer
	@mkdir -p apis
	@mkdir -p web/apis
	protoc ${PROTO_INCLUDE} ${PROTO_PATH}/checkin/protobuf/empty.proto --go_out=$(PROTO_MAPPINGS)plugins=grpc:apis --gopherjs_out=$(WEB_PROTO_MAPPINGS)plugins=grpc:web/apis
	protoc ${PROTO_INCLUDE} ${PROTO_PATH}/checkin/checkin/v1alpha/*.proto --go_out=$(PROTO_MAPPINGS)plugins=grpc:apis  --gopherjs_out=$(WEB_PROTO_MAPPINGS)plugins=grpc:web/apis
	protoc ${PROTO_INCLUDE} ${PROTO_PATH}/checkin/user/v1alpha/*.proto --go_out=$(PROTO_MAPPINGS)plugins=grpc:apis  --gopherjs_out=$(WEB_PROTO_MAPPINGS)plugins=grpc:web/apis

envcheck::
	$(call executable_check,protoc,protoc)
	$(call executable_check,protoc-gen-go,protoc-gen-go)
	$(call executable_check,protoc-gen-gopherjs,protoc-gen-gopherjs)
