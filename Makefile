build:
	@go build -o bin/agrisense_backend

run: build
	@./bin/agrisense_backend

test:
	@go test -w ./..