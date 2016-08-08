package model

//go:generate reform

//reform:area_type
type AreaType struct {
	ID          int32  `reform:"id,pk"`
	GID         string `reform:"gid"`
	Name        string `reform:"name"`
	Parent      *int32 `reform:"parent"`
	ChildOrder  int32 `reform:"child_oreder"`
	Description string `reform:"description"`
}
