.PHONY: \
	all \
	lint \
	vet \
	fmt \
	fmtcheck \
	pretest \
	test \
	gui \
	doc \
	cov \
	clean \
	build \
	install \
	generate \
	linux \
	darwin


SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
STUBS = ./srv/ctl ./srv/cmd ./srv/usr
DIST_LINUX = steam-$(STEAM_RELEASE_VERSION)-linux-amd64
DIST_DARWIN = steam-$(STEAM_RELEASE_VERSION)-darwin-amd64
WWW=./var/master/www

all: build

install: build
	go install

build:
	go build

gui:
	$(MAKE) -C gui

generate:
	go generate ./...
	$(foreach stub,$(STUBS),go fmt $(stub)/service.go || exit;)

lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

vet:
	@-go get -v golang.org/x/tools/cmd/vet
	go vet

fmt:
	gofmt -w $(SRCS)

fmtcheck:
	$(foreach file,$(SRCS),gofmt $(file) | diff -u $(file) - || exit;)

pretest: lint vet fmtcheck

test: pretest
	go test

cov:
	@ go get -v github.com/axw/gocov/gocov
	@ go get golang.org/x/tools/cmd/cover
	gocov test | gocov report

clean:
	go clean

linux:
	rm -rf ./dist/$(DIST_LINUX)
	env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=$(STEAM_RELEASE_VERSION) -X main.BUILD_DATE=`date -u +%Y-%m-%dT%H:%M:%S%z`"
	mkdir -p ./dist/$(DIST_LINUX) && mv ./steamY ./dist/$(DIST_LINUX)/steam

darwin:
	rm -rf ./dist/$(DIST_DARWIN)
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.VERSION=$(STEAM_RELEASE_VERSION) -X main.BUILD_DATE=`date -u +%Y-%m-%dT%H:%M:%S%z`"
	mkdir -p ./dist/$(DIST_DARWIN)/steam && mv ./steamY ./dist/$(DIST_DARWIN)/steam


