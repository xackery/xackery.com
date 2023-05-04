VERSION ?= 1.0.0
NAME := xackery.com

.PHONY: server
server:
	@hugo server
.PHONY: build
build: 
	@#rm -rf public/*
	@hugo -b https://xackery.com/
relogin:
	firebase logout
	firebase login
	firebase use xackery
deploy: build
	@firebase deploy
set-version:
	@echo "VERSION=${VERSION}" >> $$GITHUB_ENV