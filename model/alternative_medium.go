package model

//go:generate reform

//reform:alternative_medium
type AlternativeMedium struct {
	ID                   int32  `reform:"id,pk"`
	Name                 string `reform:"name"`
	Medium               int32 `reform:"medium"`
	AlternativeReleaseId int32 `reform:"alternative_release"`
}

// CHECK (name != '')
