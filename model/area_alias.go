package model

import "time"

//go:generate reform

//reform:area_alias
type AreaAlias struct {
	ID               int32  `reform:"id,pk"`
	Name             string `reform:"name"`
	AreaId           int32 `reform:"area"`
	Locale           string `reform:"locale"`
	EditPending      int32 `reform:"edit_pending"`
	LastUpdated      *time.Time `reform:"last_updated"`
	TypeId           int32 `reform:"type"`
	SortName         string `reform:"sort_name"`
	BeginDateYear    int16 `reform:"begin_date_year"`
	BeginDateMonth   int16 `reform:"begin_date_month"`
	BeginDateDay     int16 `reform:"begin_date_day"`
	EndDateYear      int16 `reform:"end_date_year"`
	EndDateMonth     int16 `reform:"end_date_month"`
	EndDateDay       int16 `reform:"end_date_day"`
	PrimaryForLocale bool `reform:"primary_for_locale"`
	Ended            bool `reform:"ended"`
}
