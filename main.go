package main

import (
	"context"
	"log"
	"os"

	"fmt"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"
)

func main() {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}

	toke, err := config.Token(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	httpClient := spotifyauth.New().Client(ctx, toke)
	client := spotify.New(httpClient)

	// Get a playlist
	playlist, err := client.GetPlaylist(ctx, "37i9dQZF1DXcBWIGoYBM5M")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(playlist.Name)
}
