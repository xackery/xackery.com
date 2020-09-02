.PHONY: build
build: 
	@#rm -rf public/*
	@hugo -b https://xackery.com
.PHONY: deploy
deploy: build
	@firebase deploy