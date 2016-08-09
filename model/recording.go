package model

import "time"

//go:generate reform

//reform:recording
type Recording struct {
	ID             int32  `reform:"id,pk"`
	GID            string `reform:"gid"`
	Name           string `reform:"name"`
	ArtistCreditId int32 `reform:"artist_credit"`
	Length         *int32 `reform:"length"`
	Comment        string `reform:"comment"`
	EditsPending   int32 `reform:"edits_pending"`
	LastUpdated    *time.Time        `reform:"last_updated"`
	Video          bool `reform:"video"`
}