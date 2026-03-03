package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/scheduler"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/simulation"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/storage"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func main() {

	// =========================================
	// CLI MODE: gettx <height> <index>
	// =========================================
	if len(os.Args) == 4 && os.Args[1] == "gettx" {

		height, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("invalid block height")
		}

		index, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal("invalid transaction index")
		}

		db, err := storage.Open("aegisq.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		blockObj, err := db.GetBlock(uint64(height))
		if err != nil {
			log.Fatal(err)
		}

		if index >= len(blockObj.Transactions) {
			log.Fatal("transaction index out of range")
		}

		tx := blockObj.Transactions[index]

		printTxDetails(blockObj.Index, index, tx)
		return
	}

	// =========================================
	// CLI MODE: gettxhash <hash>
	// =========================================
	if len(os.Args) == 3 && os.Args[1] == "gettxhash" {

		hash := os.Args[2]

		db, err := storage.Open("aegisq.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		blockObj, index, err := db.GetTransactionByHash(hash)
		if err != nil {
			log.Fatal(err)
		}

		tx := blockObj.Transactions[index]

		printTxDetails(blockObj.Index, index, tx)
		return
	}

	// =========================================
	// NORMAL BLOCK PRODUCTION MODE
	// =========================================

	signer, err := crypto.NewDilithiumSigner()
	if err != nil {
		panic(err)
	}
	// 1️⃣ Initialize Validators
	var validators []*identity.NodeIdentity

	for i := 1; i <= 4; i++ {
		node, err := identity.NewNodeIdentity(
			fmt.Sprintf("validator-%d", i),
			signer,
		)
		if err != nil {
			log.Fatal(err)
		}
		validators = append(validators, node)
	}

	fmt.Println("Validators initialized.")

	// 2️⃣ Governance Layer
	vs := consensus.NewValidatorSet()
	for _, v := range validators {
		vs.AddValidator(v.NodeID, v.PublicKey)
	}

	// 3️⃣ Scheduler
	sched := scheduler.NewRoundRobinScheduler(vs)

	// 4️⃣ Open Database
	db, err := storage.Open("aegisq.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	height, err := db.GetLatestHeight()
	if err != nil {
		log.Fatal(err)
	}

	var previousHash []byte

	if height > 0 {
		lastBlock, err := db.GetBlock(height)
		if err != nil {
			log.Fatal(err)
		}
		previousHash = lastBlock.Hash
		fmt.Println("Restored height:", height)
	} else {
		fmt.Println("No chain found. Starting fresh.")
	}

	// 5️⃣ Select Leader
	view := 0

	leaderID, err := sched.GetLeader(int(height+1), view)
	if err != nil {
		log.Fatal(err)
	}

	var leader *identity.NodeIdentity
	for _, v := range validators {
		if v.NodeID == leaderID {
			leader = v
			break
		}
	}

	if leader == nil {
		log.Fatal("leader not found")
	}

	fmt.Println("Leader selected:", leader.NodeID)

	// 6️⃣ Generate 10K Synthetic Transactions
	startTx := time.Now()

	txs, err := simulation.GenerateSyntheticDataset(10000, leader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated synthetic storage transactions:", len(txs))
	fmt.Println("Transaction generation time:", time.Since(startTx))

	// 7️⃣ Create & Finalize Block
	startFinalize := time.Now()

	newBlock := block.NewBlock(
		int(height+1),
		view,
		previousHash,
		txs,
	)

	if err := newBlock.Finalize(leader); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block finalize time:", time.Since(startFinalize))
	fmt.Println("Proposed block height:", newBlock.Index)

	blockHashString := fmt.Sprintf("%x", newBlock.Hash)

	// 8️⃣ BFT Voting
	votePool := consensus.NewVotePool(vs)

	// PREPARE
	for _, v := range validators {
		vote := consensus.Vote{
			ValidatorID: v.NodeID,
			BlockHash:   blockHashString,
			View:        view,
			Type:        consensus.Prepare,
		}
		_ = votePool.AddVote(vote)
	}

	if !votePool.HasQuorum(blockHashString, view, consensus.Prepare) {
		fmt.Println("Prepare quorum NOT reached.")
		return
	}

	fmt.Println("Prepare quorum reached.")

	// COMMIT
	for _, v := range validators {
		vote := consensus.Vote{
			ValidatorID: v.NodeID,
			BlockHash:   blockHashString,
			View:        view,
			Type:        consensus.Commit,
		}
		_ = votePool.AddVote(vote)
	}

	if !votePool.HasQuorum(blockHashString, view, consensus.Commit) {
		fmt.Println("Commit quorum NOT reached.")
		return
	}

	fmt.Println("Commit quorum reached.")

	// 9️⃣ Persist Block
	if err := db.SaveBlock(newBlock); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block committed at height:", newBlock.Index)

	printBlockSummary(newBlock)
}

// =========================================
// Helper Functions
// =========================================

func printTxDetails(height int, index int, tx *transaction.Transaction) {
	fmt.Println("----- Transaction Details -----")
	fmt.Println("Block Height:", height)
	fmt.Println("Transaction Index:", index)
	fmt.Println("Sender:", tx.SenderID)
	fmt.Println("Algorithm:", tx.Algorithm)
	fmt.Println("DataHash:", tx.DataHash)
	fmt.Println("Metadata:", tx.Metadata)
	fmt.Println("Timestamp:", tx.Timestamp)
	fmt.Printf("Signature: %x\n", tx.Signature)
	fmt.Println("--------------------------------")
}

func printBlockSummary(b *block.Block) {

	fmt.Println("\n========= BLOCK SUMMARY =========")
	fmt.Println("Height:", b.Index)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Printf("Previous: %x\n", b.PreviousHash)
	fmt.Println("Total Transactions:", len(b.Transactions))

	for i := 0; i < 5 && i < len(b.Transactions); i++ {
		tx := b.Transactions[i]
		fmt.Println("  Tx", i+1)
		fmt.Println("   Sender:", tx.SenderID)
		fmt.Println("   DataHash:", tx.DataHash)
	}

	if len(b.Transactions) > 5 {
		fmt.Println("  ...")
	}
}
