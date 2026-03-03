package storage

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"go.etcd.io/bbolt"
)

func uint64ToBytes(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func (db *DB) SaveBlock(b *block.Block) error {
	return db.conn.Update(func(tx *bbolt.Tx) error {

		blocks := tx.Bucket(BlocksBucket)
		hashIndex := tx.Bucket(HashIndexBucket)
		meta := tx.Bucket(MetaBucket)

		// Prevent duplicate hash
		if hashIndex.Get(b.Hash) != nil {
			return errors.New("block already exists")
		}

		data, err := json.Marshal(b)
		if err != nil {
			return err
		}

		heightKey := uint64ToBytes(uint64(b.Index))

		if err := blocks.Put(heightKey, data); err != nil {
			return err
		}

		if err := hashIndex.Put(b.Hash, heightKey); err != nil {
			return err
		}

		if err := meta.Put([]byte("latest_height"), heightKey); err != nil {
			return err
		}

		if err := meta.Put([]byte("latest_hash"), b.Hash); err != nil {
			return err
		}

		return nil
	})
}

func (db *DB) GetLatestHeight() (uint64, error) {
	var height uint64

	err := db.conn.View(func(tx *bbolt.Tx) error {
		meta := tx.Bucket(MetaBucket)

		h := meta.Get([]byte("latest_height"))
		if h == nil {
			height = 0
			return nil
		}

		height = bytesToUint64(h)
		return nil
	})

	return height, err
}

func (db *DB) GetBlock(height uint64) (*block.Block, error) {
	var result *block.Block

	err := db.conn.View(func(tx *bbolt.Tx) error {
		blocks := tx.Bucket(BlocksBucket)

		key := uint64ToBytes(height)
		data := blocks.Get(key)
		if data == nil {
			return errors.New("block not found")
		}

		var b block.Block
		if err := json.Unmarshal(data, &b); err != nil {
			return err
		}

		result = &b
		return nil
	})

	return result, err
}
