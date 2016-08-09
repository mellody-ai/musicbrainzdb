package model

import "time"

//go:generate reform

//reform:series
type Series struct {
	ID                  int32  `reform:"id,pk"`
	GID                 string `reform:"gid"`
	Name                string `reform:"name"`
	Comment             string `reform:"comment"`
	Type                *int32 `reform:"type"`
	OrderingAttributeId int32 `reform:"ordering_attribute"`
	OrderingTypeId      int32 `reform:"ordering_type"`
	EditsPending        int32 `reform:"edits_pending"`
	LastUpdated         *time.Time        `reform:"last_updated"`
}
