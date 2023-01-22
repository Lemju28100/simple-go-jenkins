.PHONY: build
build:
	@docker build . -t 2464410/simple-go-jenkins

.PHONY: run
run:
	@docker run -p 3000:3000 2464410/simple-go-jenkins