package model

//go:genrate reform

//reform:alternative_release
type AlternativeRelease struct {
	ID             int32 `reform:"id,pk"`
	GID            string `reform:"gid"`
	Release        int `reform:"release"`
	Name           string `reform:"name"`
	ArtistCreditId *int32 `reform:"artist_credit"`
	TypeId         int32 `reform:"type"`
	LanguageId     int32 `reform:"language"`
	ScriptId       int32 `reform:"script"`
	Comment        string `reform:"comment"`
}

// CHECK (name != '')