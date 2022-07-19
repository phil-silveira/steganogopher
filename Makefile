install:
	cd cmd/steganogopher && \
	go build && go install

coverage:
	go test -cover -coverprofile=coverage.out ./... && \
	go tool cover -html=coverage.out