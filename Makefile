build_proto:
	protoc \
	-I $(CURDIR)/internal/delivery/grpc/proto/ \
	--go_out=$(CURDIR)/internal/delivery/grpc/proto/ \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(CURDIR)/internal/delivery/grpc/proto/ \
	--go-grpc_opt=paths=source_relative \
	$(CURDIR)/internal/delivery/grpc/proto/app_position.proto

build:
	[ ! -d "$(CURDIR)/build" ] && mkdir $(CURDIR)/build
	go build -ldflags "-s -w" -o $(CURDIR)/build/appposition $(CURDIR)/cmd/appposition/main.go