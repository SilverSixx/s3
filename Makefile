.PHONY: run clean

run:
	@echo "Starting development servers..."
	@make -j 2 run-fe run-be

run-fe:
	@cd app && npm run dev

run-be:
	@cd api && go mod tidy && go build -o ./bin/s3 cmd/api/main.go && ./bin/s3

# Clean
clean:
	@rm -rf bin