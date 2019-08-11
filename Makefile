run: build
	@./playlists-by-tallinn

build:
	@go build -o playlists-by-tallinn ./internal

test:
	@go test ./...

clean:
	@go clean
