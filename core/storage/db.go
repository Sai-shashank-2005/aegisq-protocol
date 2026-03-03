package storage

import (
	"time"

	"go.etcd.io/bbolt"
)

var (
	MetaBucket      = []byte("meta")
	BlocksBucket    = []byte("blocks")
	HashIndexBucket = []byte("block_hash_index")
)

type DB struct {
	conn *bbolt.DB
}

func Open(path string) (*DB, error) {
	db, err := bbolt.Open(path, 0600, &bbolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		buckets := [][]byte{
			MetaBucket,
			BlocksBucket,
			HashIndexBucket,
		}

		for _, b := range buckets {
			_, err := tx.CreateBucketIfNotExists(b)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
