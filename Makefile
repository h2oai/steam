.PHONY: \
	all \
	setup \
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
DIST_LINUX = steam-$(STEAM_RELEASE_VERSION)-linux-amd64
DIST_DARWIN = steam-$(STEAM_RELEASE_VERSION)-darwin-amd64
WWW=./var/master/www
ASSETS = ./var/master/assets
JETTYRUNNER = jetty-runner-9.2.12.v20150709.jar

all: build gui setup

setup:
	mkdir -p $(ASSETS)
	cp $(SS_HOME)/$(JETTYRUNNER) $(ASSETS)/jetty-runner.jar
	cp $(SS_HOME)/makewar-extra.jar $(ASSETS)/

install: build
	go install

build:
	go build

gui:
	$(MAKE) -C gui

generate:
	cd ./srv/web && go generate && go fmt service.go
	cd ./master/db && go generate && go fmt encoding.go

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
	mkdir -p ./dist/$(DIST_LINUX)/var/master && mv ./steamY ./dist/$(DIST_LINUX)/steam
	cp -r $(WWW) ./dist/$(DIST_LINUX)/var/master/
	cp -r $(ASSETS) ./dist/$(DIST_LINUX)/var/master/	
	tar czfC ./dist/$(DIST_LINUX).tar.gz dist $(DIST_LINUX)

darwin:
	rm -rf ./dist/$(DIST_DARWIN)
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.VERSION=$(STEAM_RELEASE_VERSION) -X main.BUILD_DATE=`date -u +%Y-%m-%dT%H:%M:%S%z`"
	mkdir -p ./dist/$(DIST_DARWIN)/var/master&& mv ./steamY ./dist/$(DIST_DARWIN)/steam
	cp -r $(WWW) ./dist/$(DIST_DARWIN)/var/master/
	cp -r $(ASSETS) ./dist/$(DIST_DARWIN)/var/master/
	tar czfC ./dist/$(DIST_DARWIN).tar.gz dist $(DIST_DARWIN)


