package model

//go:generate reform

//reform:alternative_release_type
type AlternativeReleaseType struct {
	ID          int32  `reform:"id,pk"`
	GID         string `reform:"gid"`
	Name        string `reform:"name"`
	ParentId    *int32 `reform:"parent"`
	ChildOrder  int32 `reform:"child_order"`
	Description string `reform:"description"`
}
