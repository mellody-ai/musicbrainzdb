package model

//go:generate reform

//reform:alternative_track
type AlternativeTrack struct {
	ID             int32  `reform:"id,pk"`
	Name           string `reform:"name"`
	ArtistCreditId *int32 `reform:"artist_credit"`
	RefCount       int32 `reform:"ref_count"`
}

// CHECK (name != '' AND (name IS NOT NULL OR artist_credit IS NOT NULL))