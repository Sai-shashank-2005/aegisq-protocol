package main

import (
	"fmt"
	"log"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/scheduler"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/storage"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func main() {

	signer := &crypto.Ed25519Signer{}

	// ===============================
	// 1️⃣ Initialize Validators
	// ===============================
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

	// ===============================
	// 2️⃣ Validator Governance (Layer 6)
	// ===============================
	vs := consensus.NewValidatorSet()
	for _, v := range validators {
		vs.AddValidator(v.NodeID, v.PublicKey)
	}

	// ===============================
	// 3️⃣ Initialize Scheduler (Layer 7)
	// ===============================
	sched := scheduler.NewRoundRobinScheduler(vs)

	// ===============================
	// 4️⃣ Open Persistent Storage
	// ===============================
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

	// ===============================
	// 5️⃣ Determine Leader
	// ===============================
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
		log.Fatal("leader not found in validator list")
	}

	fmt.Println("Leader selected:", leader.NodeID)

	// ===============================
	// 6️⃣ Create Transaction
	// ===============================
	tx := transaction.NewTransaction(
		leader,
		fmt.Sprintf("DATA_HASH_%d", height+1),
		fmt.Sprintf("Block %d metadata", height+1),
	)

	if err := tx.SignWithIdentity(leader); err != nil {
		log.Fatal(err)
	}

	// ===============================
	// 7️⃣ Propose Block
	// ===============================
	newBlock := block.NewBlock(
		int(height+1),
		view,
		previousHash,
		[]*transaction.Transaction{tx},
	)

	if err := newBlock.Finalize(leader); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Proposed block height:", newBlock.Index)

	blockHashString := fmt.Sprintf("%x", newBlock.Hash)

	// ===============================
	// 8️⃣ BFT Voting (Layer 9)
	// ===============================
	votePool := consensus.NewVotePool(vs)

	// PREPARE
	for _, v := range validators {

		if !vs.IsAuthorized(v.NodeID, v.PublicKey) {
			continue
		}

		vote := consensus.Vote{
			ValidatorID: v.NodeID,
			BlockHash:   blockHashString,
			View:        view,
			Type:        consensus.Prepare,
		}

		if err := votePool.AddVote(vote); err != nil {
			fmt.Println("Prepare vote rejected:", err)
		}
	}

	if !votePool.HasQuorum(blockHashString, view, consensus.Prepare) {
		fmt.Println("Prepare quorum NOT reached.")
		return
	}

	fmt.Println("Prepare quorum reached.")

	// COMMIT
	for _, v := range validators {

		if !vs.IsAuthorized(v.NodeID, v.PublicKey) {
			continue
		}

		vote := consensus.Vote{
			ValidatorID: v.NodeID,
			BlockHash:   blockHashString,
			View:        view,
			Type:        consensus.Commit,
		}

		if err := votePool.AddVote(vote); err != nil {
			fmt.Println("Commit vote rejected:", err)
		}
	}

	if !votePool.HasQuorum(blockHashString, view, consensus.Commit) {
		fmt.Println("Commit quorum NOT reached.")
		return
	}

	fmt.Println("Commit quorum reached.")

	// ===============================
	// 9️⃣ Finality (Layer 10)
	// ===============================
	if err := db.SaveBlock(newBlock); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block committed at height:", newBlock.Index)

	// ===============================
	// 🔟 Print Full Chain
	// ===============================
	printFullChain(db)
}

func printFullChain(db *storage.DB) {

	height, err := db.GetLatestHeight()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n========= FULL CHAIN =========")

	var prevHash []byte

	for i := uint64(1); i <= height; i++ {

		b, err := db.GetBlock(i)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Height:", b.Index)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Previous: %x\n", b.PreviousHash)

		if i > 1 && string(b.PreviousHash) != string(prevHash) {
			fmt.Println("⚠ Chain linkage broken at height", i)
		}

		prevHash = b.Hash

		for j, tx := range b.Transactions {
			fmt.Println("  Tx", j+1)
			fmt.Println("   Sender:", tx.SenderID)
			fmt.Println("   DataHash:", tx.DataHash)
			fmt.Println("   Metadata:", tx.Metadata)
			fmt.Printf("   Signature: %x\n", tx.Signature)
		}

		fmt.Println("--------------------------------")
	}
}
