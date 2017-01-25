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
	release \
	debian_package \
	rpm_package \
	dist/rpm/centos-6 \
	dist/rpm/centos-7 

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
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
RPM_OUT_DIR = $(ROOT_DIR)/dist/rpm

all: build gui ssb launcher

install: build
	go install

build:
	go build

gui:
	cd $(GUI) && rm -rf node_modules && rm -rf typings && npm cache clean && npm install typings && npm install && ./node_modules/.bin/typings install && npm run webpack

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
	rm -rf var/master/db
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
	rm -rf tmp target
	rm -f packaging/env.list
	rm -rf dist

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

packaging/env.list:
	env | grep STEAM >> $@

dist/rpm/centos-7: packaging/env.list
	packaging/centos-7/build

dist/rpm/centos-6: packaging/env.list
	packaging/centos-6/build

release: ssb launcher linux

debian_package:
	@echo STEAM_VERSION is $(STEAM_VERSION)
	@echo STEAM_TAR_GZ is $(STEAM_TAR_GZ)
	@echo STEAM_TAR_GZ_URL is $(STEAM_TAR_GZ_URL)
	
	rm -fr tmp
	mkdir tmp
	
	rsync -a packaging/debian tmp/
	sed "s/SUBST_PACKAGE_VERSION/$(STEAM_VERSION)/" packaging/debian/steam/DEBIAN/control > tmp/debian/steam/DEBIAN/control
	pwd
	
	(cd tmp && wget $(STEAM_TAR_GZ_URL))
	pwd
	
	mkdir -p tmp/debian/steam/opt/h2oai
	pwd
	
	(cd tmp/debian/steam/opt/h2oai && tar zxvf ../../../../$(STEAM_TAR_GZ))
	(cd tmp/debian/steam/opt/h2oai && mv steam-$(STEAM_VERSION)-linux-amd64 steam)
	pwd
	
	(cd tmp/debian && dpkg-deb -b steam .)
	pwd
	
	mkdir -p target
	cp -p tmp/debian/steam_$(STEAM_VERSION)_amd64.deb target

rpm_package:
ifndef STEAM_RELEASE_VERSION
	$(error STEAM_RELEASE_VERSION is not defind)
endif

	@echo STEAM_RELEASE_VERSION is $(STEAM_RELEASE_VERSION)

	rm -fr $(RPM_OUT_DIR)
	mkdir -p $(RPM_OUT_DIR)
	
	rsync -a packaging/rpm/ $(RPM_OUT_DIR)
	pwd
	
	mkdir -p $(RPM_OUT_DIR)/steam/opt/h2oai/steam/
	cp -r dist/steam-${STEAM_RELEASE_VERSION}-linux-amd64/. $(RPM_OUT_DIR)/steam/opt/h2oai/steam/
	pwd
	
	(cd dist && echo -e "\n" | setsid fpm -s dir \
		-t rpm \
		-n steam \
		-v $(STEAM_RELEASE_VERSION) \
		--vendor H2O.ai \
		--url http://h2o.ai/download \
		--description "Steam Cluster Manager" \
		--license "AGPLv3" \
		--category "System Environment/Daemons" \
		--depends "haproxy >= 1.5, /sbin/service, /sbin/chkconfig" \
		--pre-install ../packaging/rpm/SCRIPTS/pre \
		--post-install ../packaging/rpm/SCRIPTS/post \
		--pre-uninstall ../packaging/rpm/SCRIPTS/preun \
		--post-uninstall ../packaging/rpm/SCRIPTS/postun \
		--config-files "/etc/steam/steam.conf" \
		--force \
		-C $(RPM_OUT_DIR)/steam)
	pwd

