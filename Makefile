install:
	go install ./cmd/steganogopher

build_exe:
	go build  -o ./build/ ./cmd/steganogopher

coverage:
	go test -cover -coverprofile=coverage.out ./... && \
	go tool cover -html=coverage.out
