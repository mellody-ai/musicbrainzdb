package main

import (
	"bitbucket.org/mellodycloud/mellody-musicbrainzdb"
	"os"
	"fmt"
)

var (
	CONNECTION_STRING = os.Getenv("CONNECTION_STRING")
)

func main() {
	fmt.Println("Query examples")

	context, err := mellody_musicbrainzdb.NewDBContext(CONNECTION_STRING)
	if err != nil {
		panic(err)
	}

	artist, _ := context.Artist.Get("17db8438-f314-46ea-98e3-8a116c38d504")
	fmt.Println(artist.Name)
}
