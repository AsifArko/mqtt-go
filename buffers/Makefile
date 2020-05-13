PROFILE_DIR?=profile
COMMON_DIR?=common

$(COMMON_DIR):
		protoc --gofast_out=plugins=grpc:$(GOPATH)/src $(COMMON_DIR)*.proto

$(PROFILE_DIR):
		make $(COMMON_DIR) && protoc --gofast_out=plugins=grpc:$(GOPATH)/src $(PROFILE_DIR)_*.proto