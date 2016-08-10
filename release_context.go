package mellody_musicbrainzdb

import (
	"gopkg.in/reform.v1"
	"bitbucket.org/mellodycloud/mellody-musicbrainzdb/model"
)

type ReleaseContext struct {
	DB *reform.DB
}

func NewReleaseContext(db *reform.DB) ReleaseContext {
	return ReleaseContext{db}
}

func (rc ReleaseContext) Get(gid string) (*model.Release, error) {
	if entity, err := rc.DB.FindOneFrom(model.ReleaseTable, "gid", gid); err != nil {
		return nil, err
	} else {
		release := entity.(*model.Release)
		return release, nil
	}
}