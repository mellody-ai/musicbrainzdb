package model

import "time"

//go:generate reform

//reform:medium
type Medium struct {
	ID           int32      `reform:"id,pk"`
	ReleaseId    int32      `reform:"release"`
	Position     int32      `reform:"position"`
	FormatId     *int32     `reform:"format"`
	Name         string     `reform:"name"`
	EditsPending int32      `reform:"edits_pending"`
	LastUpdated  *time.Time `reform:"last_updated"`
	TrackCount   int32      `reform:"track_count"`
}
