.PHONY: \
	all \
	lint \
	vet \
	fmt \
	fmtcheck \
	pretest \
	test \
	gui \
	guitest \
	js \
	ssb \
	launcher \
	doc \
	cov \
	clean \
	build \
	install \
	generate \
	cli-markdown \
	linux \
	darwin \
	release


SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
DIST_LINUX = steam-$(STEAM_RELEASE_VERSION)-linux-amd64
DIST_DARWIN = steam-$(STEAM_RELEASE_VERSION)-darwin-amd64
SLA=./tools/steamlauncher
SSB=./prediction-service-builder
WWW=./var/master/www
DB=./var/master/db
GUI=./gui
ASSETS = ./var/master/assets
SCRIPTS = ./scripts
JETTYRUNNER = jetty-runner-9.2.12.v20150709.jar

all: build gui ssb launcher

install: build
	go install

build:
	go build

gui:
	cd $(GUI) && rm -rf typings && rm -rf node_modules && npm cache clean && npm install && npm run webpack

guitest:
	cd $(GUI) && npm test

js:
	cd $(GUI) && npm run webpack

ssb:
	cd $(SSB) && ./gradlew build
	mkdir -p $(ASSETS)
	cp $(SSB)/$(JETTYRUNNER) $(ASSETS)/jetty-runner.jar
	cp $(SSB)/build/libs/ROOT.war $(ASSETS)/

launcher:
	cd $(SLA) && go build

generate:
	cd ./tools/piping && go build && go install
	piping
	go fmt ./srv/web/service.go
	go fmt ./cli2/cli.go
	cd ./tools/crudr && go build && go install
	cd ./master/data && go generate && go fmt scans.go && go fmt crud.go

cli-markdown:
	cd ./tools/cli-md && go build && go install
	mkdir -p ./docs/cli
	rm -f ./docs/cli/*.md
	cli-md ./docs/cli/
	mv ./docs/cli/steam.md ./docs/cli/README.md

lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

vet:
	go vet

fmt:
	gofmt -w $(SRCS)

fmtcheck:
	$(foreach file,$(SRCS),gofmt $(file) | diff -u $(file) - || exit;)

pretest: lint vet fmtcheck

test:
	cd tests && ./goh2orunner.sh

reset:
	rm -rf var/master/model
	rm -rf var/master/project

cov:
	@ go get -v github.com/axw/gocov/gocov
	@ go get golang.org/x/tools/cmd/cover
	gocov test | gocov report

clean:
	go clean
	rm -rf var
	cd $(SSB) && ./gradlew clean

linux: gui
	rm -rf ./dist/$(DIST_LINUX)
	env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=$(STEAM_RELEASE_VERSION) -X main.BUILD_DATE=`date -u +%Y-%m-%dT%H:%M:%S%z`"
	cd $(SLA) && env GOOS=linux GOARCH=amd64 go build
	mkdir -p ./dist/$(DIST_LINUX)/var/master && mv ./steam ./dist/$(DIST_LINUX)/steam
	cp LICENSE ./dist/$(DIST_LINUX)/LICENSE
	mv $(SLA)/steamlauncher ./dist/$(DIST_LINUX)/steam-launcher
	cp $(SLA)/config.toml ./dist/$(DIST_LINUX)/config.toml
	cp -r $(WWW) ./dist/$(DIST_LINUX)/var/master/
	cp -r $(ASSETS) ./dist/$(DIST_LINUX)/var/master/
	cp -r $(SCRIPTS) ./dist/$(DIST_LINUX)/var/master/
	tar czfC ./dist/$(DIST_LINUX).tar.gz dist $(DIST_LINUX)

darwin: gui
	rm -rf ./dist/$(DIST_DARWIN)
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.VERSION=$(STEAM_RELEASE_VERSION) -X main.BUILD_DATE=`date -u +%Y-%m-%dT%H:%M:%S%z`"
	cd $(SLA) && env GOOS=darwin GOARCH=amd64 go build
	mkdir -p ./dist/$(DIST_DARWIN)/var/master && mv ./steam ./dist/$(DIST_DARWIN)/steam
	cp LICENSE ./dist/$(DIST_DARWIN)/LICENSE
	mv $(SLA)/steamlauncher ./dist/$(DIST_DARWIN)/steam-launcher
	cp $(SLA)/config.toml ./dist/$(DIST_DARWIN)/config.toml
	cp -r $(WWW) ./dist/$(DIST_DARWIN)/var/master/
	cp -r $(ASSETS) ./dist/$(DIST_DARWIN)/var/master/
	cp -r $(SCRIPTS) ./dist/$(DIST_DARWIN)/var/master/
	tar czfC ./dist/$(DIST_DARWIN).tar.gz dist $(DIST_DARWIN)

release: ssb launcher linux

