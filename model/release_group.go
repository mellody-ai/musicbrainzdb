package model

import "time"

//go:generate reform

//reform:release_group
type ReleaseGroup struct {
	ID             int32      `reform:"id,pk"`
	GID            string     `reform:"gid"`
	Name           string     `reform:"name"`
	ArtistCreditId int32      `reform:"artist_credit"`
	Type           *int32     `reform:"type"`
	Comment        string     `reform:"comment"`
	EditsPending   int32      `reform:"edits_pending"`
	LastUpdated    *time.Time `reform:"last_updated"`
}
