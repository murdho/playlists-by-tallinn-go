run: build
	@./playlists-by-tallinn

build:
	@go build -o playlists-by-tallinn ./development

test:
	@go test ./...

clean:
	@go clean
