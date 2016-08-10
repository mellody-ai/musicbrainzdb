package mellody_musicbrainzdb

import (
	"gopkg.in/reform.v1"
	"bitbucket.org/mellodycloud/mellody-toolchain/musicbrainz"
)

type ArtistContext struct {
	DB *reform.DB
}

func NewArtistContext(db *reform.DB) ArtistContext {
	return ArtistContext{db}
}

func (ac ArtistContext) Get(gid string) (*musicbrainz.Artist, error) {
	if entity, err := ac.DB.FindOneFrom(musicbrainz.ArtistTable, "gid", gid); err != nil {
		return nil, err
	}else {
		artist := entity.(*musicbrainz.Artist)
		return artist, nil
	}
}

func (ac ArtistContext) FindByName(name string) ([]*musicbrainz.Artist, error) {
	if entities, err := ac.DB.FindAllFrom(musicbrainz.ArtistTable, "name", name); err != nil {
		return nil, err
	}else {
		artists := make([]*musicbrainz.Artist, len(entities))
		for i, entity := range entities {
			artists[i] = entity.(*musicbrainz.Artist)
		}

		return artists, nil
	}
}
