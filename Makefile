SHELL = /bin/bash

SERVICES = \
	microservices/api/gateway \
	microservices/svc/greeter \
	microservices/svc/echo \

.PHONY: all
all: pre proto fmt lint

.PHONY: install
install:
	make -C .devcontainer install


.PHONY: pre
pre:
	@poetry install --no-root
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
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "alice"}' localhost:55050 skeleton.greeter.v1.Greeter/Hello || true
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"msg": "hoge"}' localhost:55051 skeleton.echo.v1.Echo/Echo || true
