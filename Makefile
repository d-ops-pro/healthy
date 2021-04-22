GIT_VERSION = latest
NAMESPACE = default23
PACKAGE = otus-healthy
IMAGE=$(NAMESPACE)/$(PACKAGE):$(GIT_VERSION)

GOOS?=linux
GOARCH?=amd64

LISTEN=8080

build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
		go build ./cmd/$(PACKAGE)

docker-build: build
	docker build -t $(IMAGE) --build-arg PACKAGE=$(PACKAGE) .

docker-run: docker-build
	docker run -it --rm \
		-e LISTEN=$(LISTEN) \
 		$(IMAGE)

docker-push: docker-build
	docker push $(IMAGE)
