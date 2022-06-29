.PHONY: acceptance bench compile fuzzy racy unit

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

acceptance: ## Run acceptance tests
	@docker-compose -f ./acceptance/acceptance-tester.yml up --build --abort-on-container-exit --exit-code-from acceptance-tester

bench: ## Run benchmark
	@go test -bench=. -count 3 main.go bench_test.go

compile:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o bin/go-test-kit main.go

fuzzy: ## Run fuzzy tests
	@go test -fuzz=FuzzSum -fuzztime=20s  main.go fuzz_test.go 
	
racy: ## Run racy tests
	go test -race main.go mock_random.go racy_test.go

unit: ## Run unit tests
	@go test -v main.go unit_test.go table_test.go

coverage: ## Run coverage report
		@go test ./pkg/... ./controllers... -coverprofile cover.out
		@go tool cover -html=cover.out