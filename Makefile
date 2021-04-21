GIT_VERSION = latest
NAMESPACE = default23
PACKAGE = otus-healthy

IMAGE=$(NAMESPACE)/$(PACKAGE):$(GIT_VERSION)

PHONY: build
build: ## Build the project binary
	go build ./cmd/$(PACKAGE)

PHONY: docker
docker:
	docker build -t $(IMAGE) --build-arg PACKAGE=$(PACKAGE) .
	docker push $(IMAGE)
