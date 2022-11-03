install:
	go install

test:
	go test ./...

coverage:
	go test -cover -coverprofile=coverage.out ./... && \
	go tool cover -html=coverage.out
