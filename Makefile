# Build
.PHONY: build
build:
	@cd ./cmd && go build -o ../build/app

# Run project as dev
.PHONY: run
run:
	@cd ./cmd && go run main.go
