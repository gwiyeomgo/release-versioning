package db

import (
	"gihub.com/gwiyeomgo/release-versioning/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName        = "release-versions"
	dataBucket    = "data"
	releaseBucket = "release"
)

//(1)Singleton 패턴을 사용해서 export 되지 않는 벼누를 마들고

var db *bolt.DB

//interface 를 위한 struct
type DB struct{}

func (d DB) FindBlock(hash string) []byte {
	return findVersion(hash)
}
func (d DB) Create(hash string, data []byte) {
	create(hash, data)
}

func InitDB() {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(releaseBucket))
			return err
		})
		utils.HandleErr(err)
	}
}

func Close() {
	db.Close()
}
func create(key string, data []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(releaseBucket))
		err := bucket.Put([]byte(key), data)
		return err
	})
	utils.HandleErr(err)
}

func findVersion(hash string) []byte {
	var data []byte
	//DB에 blocksBucket 에서 특정 블록을 찾는다
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(releaseBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
