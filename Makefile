.PHONY: goimports
goimports:
	cd /tmp && go get golang.org/x/tools/cmd/goimports

.PHONY: server_generate
server_generate:
	./.bin/server_generator_mod ./api

.PHONY: run_server
run_server:
	PORT=8080 go run cmd/{main.go,init.go}