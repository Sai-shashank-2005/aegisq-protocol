package main

import (
	"fmt"
	"log"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/block"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/consensus"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/storage"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

func main() {

	signer := &crypto.Ed25519Signer{}

	// ===============================
	// 1️⃣ Create 4 Validators
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
	// 2️⃣ Build Validator Set (Layer 6)
	// ===============================
	vs := consensus.NewValidatorSet()

	for _, v := range validators {
		vs.AddValidator(v.NodeID, v.PublicKey)
	}

	votePool := consensus.NewVotePool(vs)

	// ===============================
	// 3️⃣ Open Database
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
	// 4️⃣ Deterministic Leader (Layer 7)
	// ===============================
	view := 0
	leader := validators[view%len(validators)]

	fmt.Println("Leader selected:", leader.NodeID)

	// ===============================
	// 5️⃣ Create Transaction
	// ===============================
	tx := transaction.NewTransaction(
		leader,
		fmt.Sprintf("DATA_HASH_%d", height+1),
		fmt.Sprintf("Block %d metadata", height+1),
	)

	err = tx.SignWithIdentity(leader)
	if err != nil {
		log.Fatal(err)
	}

	// ===============================
	// 6️⃣ Create Proposed Block
	// ===============================
	newBlock := block.NewBlock(
		int(height+1),
		view,
		previousHash,
		[]*transaction.Transaction{tx},
	)

	err = newBlock.Finalize(leader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Proposed block height:", newBlock.Index)

	blockHashString := fmt.Sprintf("%x", newBlock.Hash)

	// ===============================
	// 7️⃣ PREPARE PHASE (Layer 9)
	// ===============================
	for _, v := range validators {

		// Authorization check (Layer 6)
		if !vs.IsAuthorized(v.NodeID, v.PublicKey) {
			fmt.Println("Unauthorized validator:", v.NodeID)
			continue
		}

		vote := consensus.Vote{
			ValidatorID: v.NodeID,
			BlockHash:   blockHashString,
			View:        view,
			Type:        consensus.Prepare,
		}

		err := votePool.AddVote(vote)
		if err != nil {
			fmt.Println("Prepare vote rejected:", err)
		}
	}

	if !votePool.HasQuorum(blockHashString, view, consensus.Prepare) {
		fmt.Println("Prepare quorum NOT reached.")
		return
	}

	fmt.Println("Prepare quorum reached.")

	// ===============================
	// 8️⃣ COMMIT PHASE (Layer 9)
	// ===============================
	for _, v := range validators {

		if !vs.IsAuthorized(v.NodeID, v.PublicKey) {
			fmt.Println("Unauthorized validator:", v.NodeID)
			continue
		}

		vote := consensus.Vote{
			ValidatorID: v.NodeID,
			BlockHash:   blockHashString,
			View:        view,
			Type:        consensus.Commit,
		}

		err := votePool.AddVote(vote)
		if err != nil {
			fmt.Println("Commit vote rejected:", err)
		}
	}

	if !votePool.HasQuorum(blockHashString, view, consensus.Commit) {
		fmt.Println("Commit quorum NOT reached.")
		return
	}

	fmt.Println("Commit quorum reached.")

	// ===============================
	// 9️⃣ Persist Block (Layer 10)
	// ===============================
	err = db.SaveBlock(newBlock)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block committed at height:", newBlock.Index)
}
