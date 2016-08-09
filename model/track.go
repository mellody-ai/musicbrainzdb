package model

import "time"

//go:generate reform

//reform:track
type Track struct {
	ID             int32      `reform:"id,pk"`
	GID            string     `reform:"gid"`
	RecordingId    int32      `reform:"recording"`
	MediumId       int32      `reform:"medium"`
	Position       int32      `reform:"position"`
	Number         string     `reform:"number"`
	Name           string     `reform:"name"`
	ArtistCreditId int32      `reform:"artist_credit"`
	Length         int32      `reform:"length"`
	EditsPending   int32      `reform:"edits_pending"`
	LastUpdated    *time.Time `reform:"last_updated"`
	IsDataTrack    bool       `reform:"is_data_track"`
}
