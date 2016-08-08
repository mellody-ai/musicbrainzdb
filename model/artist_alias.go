package model

import "time"

//go:generate reform

//reform:artist_alias
type ArtistAlias struct {
	ID             int32  `reform:"id,pk"`
	ArtistId       int32 `reform:"artist"`
	Name           string `reform:"name"`
	Locale         string `reform:"locale"`
	SortName       string `reform:"sort_name"`
	BeginDateYear  int16 `reform:"begin_date_year"`
	BeginDateMonth int16 `reform:"begin_date_month"`
	BeginDateDay   int16 `reform:"begin_date_day"`
	EndDateYear    int16 `reform:"end_date_year"`
	EndDateMonth   int16 `reform:"end_date_month"`
	EndDateDay     int16 `reform:"end_date_day"`
	TypeId         *int32    `reform:"type"`
	EditsPending   int32 `reform:"edits_pending"`
	LastUpdated    *time.Time        `reform:"last_updated"`
	Ended          bool `reform:"ended"`
}
