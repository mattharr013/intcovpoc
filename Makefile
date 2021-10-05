
deps:
	go install github.com/ory/go-acc@latest

coverage: FORCE
	go-acc ./...
	go tool cover -html=coverage.txt

%/coverage.txt:
	go-acc ./$*/... ./tests/integration
	go tool cover -html=coverage.txt

FORCE:
