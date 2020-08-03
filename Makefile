# just need modify this
api_protos := $(wildcard protoc/sub_protoc/*.proto)

types_protos := $(wildcard protoc/*.proto protoc/common_protoc/*.proto)
types_protos += $(api_protos)

go_pb_objs := $(patsubst %.proto, %.pb.go, $(types_protos))  	
go_micro_objs := $(patsubst %.proto, %.pb.micro.go, $(api_protos)) 
go_gw_objs := $(patsubst %.proto, %.pb.gw.go, $(api_protos))

.PHONY: tools
tools:
	go get -u -v github.com/golang/protobuf/protoc-gen-go@v1.3.0
	go get -u -v github.com/micro/micro/v2/cmd/protoc-gen-micro@master
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.4
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.14.4

.PHONY: mkdirgo go
mkdirgo:
	mkdir -p go_gens

.PHONY: mkdirgo_gens_gw go
mkdirgo_gens_gw:
	mkdir -p go_gens_gw

.PHONY: go
go: mkdirgo $(go_pb_objs) $(go_micro_objs) 	
	@echo "warning! proto-gen-go must be version 1.3.0 (use make tools first to make sure protoc-gen-go version correct)"

gw: mkdirgo_gens_gw $(go_pb_objs) $(go_gw_objs)
	cp ./go_gens/admin/*.pb.go ./go_gens_gw/admin/
	cp ./go_gens/app/*.pb.go ./go_gens_gw/app/
	cp ./go_gens/teaching/*.pb.go ./go_gens_gw/teaching/
	@echo "warning! proto-gen-go must be version 1.3.0 (use make tools first to make sure protoc-gen-go version correct)"

.PHONY: swift
swift: $(patsubst %.proto, %.pb.swift, $(types_protos))

.PHONY: js
js: $(patsubst %.proto, %.pb.js, $(types_protos))

clean:
	@rm -rf swift_gens
	@rm -rf go_gens
	@rm -rf go_gens_gw
	@rm -rf js_gens
	@echo "clean finished"

# single file build cmd
%.pb.go: %.proto
	protoc -Iprotoc -Ivendors --go_out=plugins=grpc,paths=source_relative:./go_gens $<

%.pb.gw.go: %.proto
	protoc -Iprotoc -Ivendors --grpc-gateway_out=logtostderr=true,paths=source_relative,grpc_api_configuration=$*@restful.yaml:./go_gens_gw $*.proto
	protoc -Iprotoc -Ivendors --swagger_out=logtostderr=true,grpc_api_configuration=$*@restful.yaml:./go_gens_gw $*.proto

%.pb.micro.go: %.proto
	protoc -Iprotoc -Ivendors --micro_out=.,paths=source_relative:./go_gens $<

%.pb.swift: %.proto
	@mkdir -p swift_gens
	protoc \
	-Iprotoc \
	-Ivendors \
	--swift_out=swift_gens $<

%.pb.js: %.proto
	@mkdir -p js_gens
	pbjs -r $(notdir $*) -t static-module -w es6 --es6 --keep-case -o js_gens/$(notdir $@) $< -p protoc -p vendors
	pbts -o js_gens/$(notdir $*).pb.d.ts js_gens/$(notdir $@)

