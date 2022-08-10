package db

import (
	"encoding/json"
	"fmt"
	"gihub.com/gwiyeomgo/release-versioning/utils"
	bolt "go.etcd.io/bbolt"
	"strconv"
)

const (
	dbName        = "release-versions"
	releaseBucket = "release"
)

//(1)Singleton 패턴을 사용해서 export 되지 않는 벼누를 마들고

var db *bolt.DB

//interface 를 위한 struct
type DB struct{}

func (d DB) FindLastVersion() string {
	return findLastVersion()
}

func (d DB) Create() {
	create()
}

func InitDB() {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err = tx.CreateBucketIfNotExists([]byte(releaseBucket))
			return err
		})
		utils.HandleErr(err)
	}
}

func Close() {
	db.Close()
}

func create() {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(releaseBucket))
		id, _ := b.NextSequence()
		r := struct {
			ID string
		}{}
		r.ID = strconv.FormatUint(id, 10)
		buf, err := json.Marshal(r)
		if err != nil {
			return err
		}
		fmt.Printf("%s , %s\n", r.ID, buf)
		return b.Put([]byte(r.ID), buf)
	})
	utils.HandleErr(err)
}

func findLastVersion() (version string) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(releaseBucket))
		version = strconv.FormatUint(b.Sequence(), 10)
		return nil
	})
	return version
}
