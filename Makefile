.PHONY: build
build: 
	@#rm -rf public/*
	@hugo -b https://xackery.firebaseapp.com
.PHONY: deploy
deploy: build
	@firebase deploy