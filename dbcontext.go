package mellody_musicbrainzdb

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type DBContext struct {
	DB     *reform.DB
	Artist ArtistContext
}

var context *DBContext = nil

func NewDBContext(connectionString string) (*DBContext, error) {
	if context != nil {
		return context, nil
	}

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Use new *log.Logger for logging.
	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	db := reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))
	context = &DBContext{DB: db, Artist: NewArtistContext(db)}

	return context, nil
}
