package storage

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"go.etcd.io/bbolt"
)

var (
	MetaBucket      = []byte("meta")
	BlocksBucket    = []byte("blocks")
	HashIndexBucket = []byte("block_hash_index")
	TxIndexBucket   = []byte("tx_index")
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
			TxIndexBucket,
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

func uint64ToBytes(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

//
// ==============================
// SAVE BLOCK
// ==============================
//

func (db *DB) SaveBlock(b *block.Block) error {

	return db.conn.Update(func(tx *bbolt.Tx) error {

		blocks := tx.Bucket(BlocksBucket)
		hashIndex := tx.Bucket(HashIndexBucket)
		meta := tx.Bucket(MetaBucket)
		txIndex := tx.Bucket(TxIndexBucket)

		// Prevent duplicate block
		if hashIndex.Get(b.Hash) != nil {
			return errors.New("block already exists")
		}

		data, err := json.Marshal(b)
		if err != nil {
			return err
		}

		heightKey := uint64ToBytes(uint64(b.Index))

		// Store block by height
		if err := blocks.Put(heightKey, data); err != nil {
			return err
		}

		// Index block hash → height
		if err := hashIndex.Put(b.Hash, heightKey); err != nil {
			return err
		}

		// Index transactions
		for i, txObj := range b.Transactions {

			txKey := []byte(txObj.DataHash)

			indexData := struct {
				Height uint64
				Index  int
			}{
				Height: uint64(b.Index),
				Index:  i,
			}

			indexBytes, err := json.Marshal(indexData)
			if err != nil {
				return err
			}

			if err := txIndex.Put(txKey, indexBytes); err != nil {
				return err
			}
		}

		// Update metadata
		if err := meta.Put([]byte("latest_height"), heightKey); err != nil {
			return err
		}

		if err := meta.Put([]byte("latest_hash"), b.Hash); err != nil {
			return err
		}

		return nil
	})
}

//
// ==============================
// READ BLOCK
// ==============================
//

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

//
// ==============================
// O(1) TX LOOKUP
// ==============================
//

func (db *DB) GetTransactionByHash(hash string) (*block.Block, int, error) {

	var height uint64
	var index int

	err := db.conn.View(func(tx *bbolt.Tx) error {

		txIndex := tx.Bucket(TxIndexBucket)

		data := txIndex.Get([]byte(hash))
		if data == nil {
			return errors.New("transaction not found")
		}

		var entry struct {
			Height uint64
			Index  int
		}

		if err := json.Unmarshal(data, &entry); err != nil {
			return err
		}

		height = entry.Height
		index = entry.Index

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	blockObj, err := db.GetBlock(height)
	if err != nil {
		return nil, 0, err
	}

	return blockObj, index, nil
}
