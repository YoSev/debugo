test:
	go test -race -v ./...

test+coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./... ./...

changelog:
	auto-changelog --output CHANGELOG.md

lint:
	golangci-lint run --timeout=5m ./...

test+coverage+html:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=coverage.txt
	rm coverage.txt