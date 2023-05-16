package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("SPOTIFY_CLIENT_ID"))
}
