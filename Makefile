PROJECT_NAME=clarchgo
PROJECT_ORG=sellsuki
VERSION=latest
GO_CMD=go
REPO=$(PROJECT_ORG)/$(PROJECT_NAME)

.PHONY: run

run:
	$(GO_CMD) run .
build:
	$(GO_CMD) build -o ./dist/$(BINARY_NAME) -a -v
clean:
	$(GO_CMD) clean
	rm -rf ./dist
test:
	$(GO_CMD) test -v ./...
test-integrate: test-integrate-repo test-integrate-use-case

test-integrate-repo:
	$(GO_CMD) test -v ./...
test-integrate-use-case:
	$(GO_CMD) test -v ./...

build-docker:
	docker build -t $(REPO):$(VERSION) .
push-docker:
	docker push $(REPO):$(VERSION)
test-docker:
	docker build --no-cache --rm --target=tester .
