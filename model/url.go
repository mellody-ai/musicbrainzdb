package model

import "time"

//go:generate reform

//reform:url
type Url struct {
	ID           int32  `reform:"id,pk"`
	GID          string `reform:"gid"`
	Url          string `reform:"url"`
	EditsPending int32 `reform:"edits_pending"`
	LastUpdated  *time.Time        `reform:"last_updated"`
}