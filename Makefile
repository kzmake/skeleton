SHELL = /bin/bash

SERVICES = \
	backend/api/gateway \
	backend/svc/greeter \
	backend/svc/echo \

.PHONY: all
all: pre proto fmt lint

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


.PHONY: build
build:
	skaffold build


.PHONY: kind
kind:
	kind get clusters -q | grep "skeleton" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name skeleton

.PHONY: dev
dev:
	skaffold dev


.PHONY: deploy-production
deploy-production:
	docker login ghcr.io
	skaffold run -p production

.PHONY: destroy-production
destroy-production:
	skaffold delete -p production

.PHONY: http
http:
	curl -i localhost:58080/greeter/v1/hello -H "Content-Type: application/json" -d '{"name": "alice"}'
	curl -i localhost:58080/echo/v1/echo -H "Content-Type: application/json" -d '{"msg": "hoge"}'

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext -rpc-header 'dapr-app-id: svc-greeter' -d '{"name": "alice"}' localhost:50001 skeleton.greeter.v1.Greeter/Hello || true
	grpcurl -protoset <(buf build -o -) -plaintext -rpc-header 'dapr-app-id: svc-echo' -d '{"msg": "hoge"}' localhost:50001 skeleton.echo.v1.Echo/Echo || true
