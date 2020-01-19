run: build
	@./playlists-by-tallinn

build:
	@go build -o playlists-by-tallinn ./cmd/playlists-by-tallinn

test:
	@go test ./...

cov:
	@go test ./... -coverprofile cover.out
	@go tool cover -html=cover.out

generate:
	@go generate ./...

tidy:
	@go mod tidy

clean:
	@go clean -i
	@rm -rf ./playlists-by-tallinn

tools:
	@go get github.com/matryer/moq
