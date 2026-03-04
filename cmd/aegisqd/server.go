package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/storage"
)

func startServer(db *storage.DB) {

	mux := http.NewServeMux()

	// ---------------------------
	// STATUS
	// ---------------------------
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		height, err := db.GetLatestHeight()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "running",
			"height": height,
		})
	})

	// ---------------------------
	// BLOCK LIST
	// ---------------------------
	mux.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		height, err := db.GetLatestHeight()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		blocks := []map[string]interface{}{}

		for i := height; i >= 1 && len(blocks) < 20; i-- {

			b, err := db.GetBlock(i)
			if err != nil {
				continue
			}

			blocks = append(blocks, map[string]interface{}{
				"height": b.Index,
				"hash":   fmt.Sprintf("%x", b.Hash),
				"txs":    len(b.Transactions),
			})
		}

		json.NewEncoder(w).Encode(blocks)
	})

	// ---------------------------
	// SINGLE BLOCK
	// ---------------------------
	mux.HandleFunc("/block/", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		heightStr := strings.TrimPrefix(r.URL.Path, "/block/")
		height, err := strconv.Atoi(heightStr)

		if err != nil {
			http.Error(w, "invalid height", 400)
			return
		}

		block, err := db.GetBlock(uint64(height))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(block)
	})

	// ---------------------------
	// TRANSACTION BY HEIGHT/INDEX
	// ---------------------------
	mux.HandleFunc("/tx/", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tx/"), "/")

		if len(parts) != 2 {
			http.Error(w, "usage: /tx/{height}/{index}", 400)
			return
		}

		height, _ := strconv.Atoi(parts[0])
		index, _ := strconv.Atoi(parts[1])

		block, err := db.GetBlock(uint64(height))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if index >= len(block.Transactions) {
			http.Error(w, "transaction index out of range", 400)
			return
		}

		json.NewEncoder(w).Encode(block.Transactions[index])
	})

	// ---------------------------
	// TRANSACTION BY HASH
	// ---------------------------
	mux.HandleFunc("/txhash/", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		hash := strings.TrimPrefix(r.URL.Path, "/txhash/")

		block, index, err := db.GetTransactionByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		tx := block.Transactions[index]

		json.NewEncoder(w).Encode(map[string]interface{}{
			"block_height": block.Index,
			"tx_index":     index,
			"transaction":  tx,
		})
	})

	fmt.Println("API server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
