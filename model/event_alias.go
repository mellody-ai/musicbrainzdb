package model

import "time"

//go:generate reform

//reform:event_alias
type EventAlias struct {
	ID               int32  `reform:"id,pk"`
	EventId          int32 `reform:"event"`
	Name             string `reform:"name"`
	BeginDateYear    int16 `reform:"begin_date_year"`
	BeginDateMonth   int16 `reform:"begin_date_month"`
	BeginDateDay     int16 `reform:"begin_date_day"`
	EndDateYear      int16 `reform:"end_date_year"`
	EndDateMonth     int16 `reform:"end_date_month"`
	EndDateDay       int16 `reform:"end_date_day"`
	Locale           string `reform:"locale"`
	SortName         string `reform:"sort_name"`
	TypeId           *int32 `reform:"type"`
	EditsPending     int32 `reform:"edits_pending"`
	LastUpdated      *time.Time        `reform:"last_updated"`
	PrimaryForLocale bool `reform:"primary_for_locale"`
	Ended            bool `reform:"ended"`
}
