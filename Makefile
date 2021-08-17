SHELL = /bin/bash

SERVICES = \
	microservices/api/gateway \
	microservices/svc/greeter \

.PHONY: all
all: proto pre fmt lint

.PHONY: install
install:
	make -C .devcontainer install


.PHONY: pre
pre:
	@for f in $(SERVICES); do make -C $$f pre; done

.PHONY: fmt
fmt:
	@for f in $(SERVICES); do make -C $$f fmt; done

.PHONY: lint
lint:
	@for f in $(SERVICES); do make -C $$f lint; done


.PHONY: proto
proto:
	buf mod update
	buf generate


.PHONY: kind
kind:
	kind get clusters -q | grep "skeleton" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name skeleton

.PHONY: dev
dev:
	skaffold run -f skaffold.dapr.yaml
	skaffold dev


.PHONY: http
http:
	curl -i localhost:58080/greeter/hello -H "Content-Type: application/json" -d '{"name": "alice"}'

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "alice"}' localhost:55050 skeleton.greeter.v1.Greeter/Hello || true
