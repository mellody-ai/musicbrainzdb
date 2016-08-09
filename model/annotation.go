package model

import "time"

//go:generate reform

//reform:annotation
type Annotation struct {
	ID        int32      `reform:"id,pk"`
	Editor    int32      `reform:"editor"`
	Text      string     `reform:"text"`
	Changelog string     `reform:"changelog"`
	Created   *time.Time `reform:"created"`
}
