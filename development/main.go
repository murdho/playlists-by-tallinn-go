package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	. "github.com/murdho/playlists-by-tallinn"
	"log"
)

func main() {
	if err := PlaylistsByTallinn(context.Background(), PubSubMessage{}); err != nil {
		log.Fatal(err)
	}
}
