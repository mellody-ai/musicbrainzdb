package model

import "time"

//go:generate reform

//reform:area
type Area struct {
	ID             int32  `reform:"id,pk"`
	GID            string `reform:"gid"`
	Name           string `reform:"name"`
	TypeId         *int    `reform:"type"`
	EditsPending   int32 `reform:"edits_pending"`
	LastUpdated    *time.Time `reform:"last_updated"`
	BeginDateYear  int16 `reform:"begin_date_year"`
	BeginDateMonth int16 `reform:"begin_date_month"`
	BeginDateDay   int16 `reform:"begin_date_day"`
	EndDateYear    int16 `reform:"end_date_year"`
	EndDateMonth   int16 `reform:"end_date_month"`
	EndDateDay     int16 `reform:"end_date_day"`
	Ended          bool `reform:"ended"`
	/*
	CHECK (
		(
		  -- If any end date fields are not null, then ended must be true
		  (end_date_year IS NOT NULL OR
		   end_date_month IS NOT NULL OR
		   end_date_day IS NOT NULL) AND
		  ended = TRUE
		) OR (
		  -- Otherwise, all end date fields must be null
		  (end_date_year IS NULL AND
		   end_date_month IS NULL AND
		   end_date_day IS NULL)
		)
	      ),
 	*/
	Comment        string `reform:"comment"`
}

