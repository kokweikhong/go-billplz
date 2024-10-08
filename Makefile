web-build:
	@echo "Building web..."
	@go build -o bin/web cmd/web/main.go

web-run: web-build
	@echo "Running web..."
	@./bin/web
