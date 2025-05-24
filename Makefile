test:
	go test -race -v ./...
	
changelog:
	auto-changelog --output CHANGELOG.md

