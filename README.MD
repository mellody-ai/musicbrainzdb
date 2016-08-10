#musicbrainzdb

It's wrapper for work with [MusicBrainz Database](https://musicbrainz.org/doc/MusicBrainz_Database/Schema)

## Dependencies ##

* [Golang](https://golang.org/) (version 1.6 or above)
* [pq](http://github.com/lib/pq)
* [Reform ORM](https://github.com/go-reform/reform)

```sh
go get gopkg.in/reform.v1
```

## Install ##

* go get *bitbucket.org/mellodycloud/mellody-musicbrainzdb.git*

## How use ##

1. import "bitbucket.org/mellodycloud/mellody-musicbrainzdb"
2. Set Env variable *CONNECTION_STRING* to connect a DB.

```sh
export CONNECTION_STRING = postgres://usr:!QAZxsw2@plar.global:5432/musicbrainz_db?sslmode=disable
```

```go
var(
	CONNECTION_STRING = os.Getenv("CONNECTION_STRING")
)
```

3. Make context

```go
context, _ := mellody_musicbrainzdb.NewDBContext(CONNECTION_STRING)
```

4. Do query

```go
artist, _ := context.Artist.Get("17db8438-f314-46ea-98e3-8a116c38d504")
fmt.Println(artist.Name)
```

More amount of examples you can find [here](https://github.com/mellody-cloud/musicbrainzdb/tree/master/examples).

## For contributors ##

### Database scheme ###

![Musicbrainz database scheme](http://wiki.musicbrainz.org/-/images/5/52/ngs.png)