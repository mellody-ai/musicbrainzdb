package model

import "time"

//go:generate reform

//reform:release
type Release struct {
	ID             int32      `reform:"id,pk"`
	GID            string     `reform:"gid"`
	Name           string     `reform:"name"`
	ArtistCreditId int32      `reform:"artist_credit"`
	ReleaseGroupId int32      `reform:"release_group"`
	StatusId       *int32     `reform:"status"`
	PackagingId    *int32     `reform:"packaging"`
	LanguageId     *int32     `reform:"language"`
	ScriptId       *int32     `reform:"script"`
	Barcode        *string    `reform:"barcode"`
	Comment        string     `reform:"comment"`
	EditsPending   int32      `reform:"edits_pending"`
	Quality        int16      `reform:"quality"`
	LastUpdated    *time.Time `reform:"last_updated"`
}
