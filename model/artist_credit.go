package model

import "time"

//go:generate reform

//reform:artist_credit
type ArtistCredit struct {
	ID          int32  `reform:"id,pk"`
	Name        string `reform:"name"`
	ArtistCount int16 `reform:"artist_count"`
	RefCount    *int32 `reform:"ref_count"`
	Created     *time.Time `reform:"created"`
}
