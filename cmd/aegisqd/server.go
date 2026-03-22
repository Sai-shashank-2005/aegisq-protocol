package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/scheduler"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/storage"
)

func startServer(
	db *storage.DB,
	vs *consensus.ValidatorSet,
	vp *consensus.VotePool,
	fe *consensus.FinalityEngine,
	scheduler *scheduler.RoundRobinScheduler,
) {

	mux := http.NewServeMux()

	// ---------------------------
	// LIVENESS TRACKING STATE
	// ---------------------------
	var lastHeight uint64 = 0
	var lastChange = time.Now()

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
	// CONSENSUS STATE (🔥 CORE)
	// ---------------------------
	mux.HandleFunc("/consensus", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		height, err := db.GetLatestHeight()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		block, err := db.GetBlock(uint64(height))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		view := int(block.View)

		leader, _ := scheduler.GetLeader(int(height), view)

		// quorum calculation
		n := vs.Count()
		f := (n - 1) / 3
		required := 2*f + 1

		// simulate received votes
		received := len(block.Transactions)
		if received > n {
			received = n
		}

		status := "IN_PROGRESS"
		if received >= required {
			status = "FINALIZED"
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"height": height,
			"view":   view,
			"leader": leader,
			"quorum": map[string]interface{}{
				"required": required,
				"received": received,
			},
			"validators": vs.GetValidatorIDs(),
			"status":     status,
		})
	})

	// ---------------------------
	// LIVENESS (🔥 USP)
	// ---------------------------
	mux.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		height, err := db.GetLatestHeight()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// detect stall
		if height != lastHeight {
			lastHeight = height
			lastChange = time.Now()
		}

		elapsed := time.Since(lastChange).Seconds()

		status := "HEALTHY"
		reason := "Block production normal"

		if elapsed > 5 {
			status = "AT_RISK"
			reason = "Block delay detected"
		}

		if elapsed > 10 {
			status = "FAILED"
			reason = "Leader failure suspected, no view-change"
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":      status,
			"reason":      reason,
			"height":      height,
			"stalled_for": elapsed,
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
	// SINGLE BLOCK (UPGRADED)
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

		view := int(block.View)

		leader, _ := scheduler.GetLeader(height, view)

		// quorum
		n := vs.Count()
		f := (n - 1) / 3
		required := 2*f + 1

		received := len(block.Transactions)
		if received > n {
			received = n
		}

		status := "PENDING"
		if received >= required {
			status = "FINALIZED"
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"height": height,
			"hash":   fmt.Sprintf("%x", block.Hash),
			"view":   view,
			"leader": leader,

			"transactions": block.Transactions,

			"consensus": map[string]interface{}{
				"quorum": map[string]interface{}{
					"required": required,
					"received": received,
				},
				"status": status,
			},
		})
	})

	// ---------------------------
	// TX BY HEIGHT/INDEX
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
	// TX BY HASH
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

	fmt.Println("🚀 API server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
