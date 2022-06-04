GOLANG_VERSION="1.18"


# Help
.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-init-module: ## Initialize go project go mod <project-name>
	read PROJECT_NAME && \
	rm -rf go.mod && \
	docker run --rm -it \
	-v $$(pwd):/data \
	-w /data/ --entrypoint="" \
	--user 1000:1000 \
	golang:$(GOLANG_VERSION) go mod init $$PROJECT_NAME


go-run-main: ## Run main.go file
	docker run --rm -it \
	-v $$(pwd):/data \
	-w /data/ --entrypoint="" \
	golang:$(GOLANG_VERSION) go run main.go


#checkov: ## checkov testes
# docker run --rm -it -v $$(pwd):/app -w /app/ bridgecrew/checkov:2.0.1190 -d /app