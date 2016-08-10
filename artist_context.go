package mellody_musicbrainzdb

import (
	"gopkg.in/reform.v1"
	"bitbucket.org/mellodycloud/mellody-musicbrainzdb/model"
)

type ArtistContext struct {
	DB *reform.DB
}

func NewArtistContext(db *reform.DB) ArtistContext {
	return ArtistContext{db}
}

func (ac ArtistContext) Get(gid string) (*model.Artist, error) {
	if entity, err := ac.DB.FindOneFrom(model.ArtistTable, "gid", gid); err != nil {
		return nil, err
	}else {
		artist := entity.(*model.Artist)
		return artist, nil
	}
}

func (ac ArtistContext) FindByName(name string) ([]*model.Artist, error) {
	if entities, err := ac.DB.FindAllFrom(model.ArtistTable, "name", name); err != nil {
		return nil, err
	}else {
		artists := make([]*model.Artist, len(entities))
		for i, entity := range entities {
			artists[i] = entity.(*model.Artist)
		}

		return artists, nil
	}
}

//func (ac ArtistContext) Releases() ([]*model.Release) {
//	ac.DB.
//}
