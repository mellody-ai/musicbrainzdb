package model

import "time"

//go:generate reform

//reform:artist
type Artist struct {
	ID             int32  `reform:"id,pk"`
	GID            string `reform:"gid"`
	Name           string `reform:"name"`
	SortName       string `reform:"sort_name"`
	BeginDateYear  int16 `reform:"begin_date_year"`
	BeginDateMonth int16 `reform:"begin_date_month"`
	BeginDateDay   int16 `reform:"begin_date_day"`
	EndDateYear    int16 `reform:"end_date_year"`
	EndDateMonth   int16 `reform:"end_date_month"`
	EndDateDay     int16 `reform:"end_date_day"`
	TypeId         *int32    `reform:"type"`
	AreaId         *int32   `reform:"area"`
	GenderId       *int32   `reform:"gender"`
	Comment        string `reform:"comment"`
	EditsPending   int32 `reform:"edits_pending"`
	LastUpdated    *time.Time        `reform:"last_updated"`
	Ended          bool `reform:"ended"`
	BeginAreaId    *int32 `reform:"begin_area"`
	EndArea        *int32 `reform:"end_area"`
}