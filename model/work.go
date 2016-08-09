package model

import "time"

//go:generate reform

//reform:work
type Work struct {
	ID           int32  `reform:"id,pk"`
	GID          string `reform:"gid"`
	Name         string `reform:"name"`
	Type         *int32 `reform:"type"`
	Comment      string `reform:"comment"`
	EditsPending int32 `reform:"edits_pending"`
	LastUpdated  *time.Time        `reform:"last_updated"`
	LanguageId   *int32 `reform:"language"`
}
