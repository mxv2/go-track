#!make

lint:
	@golint -set_exit_status ./...

.PHONY: lint
