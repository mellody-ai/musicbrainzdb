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

	fmt.Println(context.Artist.Get("c06aa285-520e-40c0-b776-83d2c9e8a6d1"))
}
