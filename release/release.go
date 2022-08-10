package release

import (
	"gihub.com/gwiyeomgo/release-versioning/db"
)

type Release struct {
	ID string `json:"id"`
}

type storage interface {
	FindLastVersion() string
	Create()
}

var dbStorage storage = db.DB{}

func FindLastVersion() (*Release, error) {

	version := dbStorage.FindLastVersion()
	r := Release{
		ID: version,
	}

	return &r, nil
}

func CreateVersion() {
	dbStorage.Create()
}
