.PHONY: server
server:
	@hugo server
.PHONY: build
build: 
	@#rm -rf public/*
	@hugo -b https://xackery.com
relogin:
	firebase logout
	firebase login
	firebase use xackery
deploy: build
	@firebase deploy