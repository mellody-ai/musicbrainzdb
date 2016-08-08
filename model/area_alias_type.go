package model

//go:generate reform

//reform:area_alias_type
type AreaAliasType struct {
	ID          int32  `reform:"id,pk"`
	GID         string `reform:"gid"`
	Name        string `reform:"name"`
	Parent      *int32 `reform:"parent"`
	ChildOrder  int32 `reform:"child_order"`
	Description string `reform:"description"`
}
