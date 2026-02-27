package block

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Sai-shashank-2005/aegisq-protocol/core/crypto"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/identity"
	"github.com/Sai-shashank-2005/aegisq-protocol/core/transaction"
)

type Block struct {
	Index        int
	Timestamp    int64
	PreviousHash []byte
	MerkleRoot   []byte
	Transactions []*transaction.Transaction
	Hash         []byte
	Validator    string
	Signature    []byte
}

func NewBlock(index int, prevHash []byte, txs []*transaction.Transaction) *Block {
	return &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		PreviousHash: prevHash,
		Transactions: txs,
	}
}

func (b *Block) computeBlockHash() ([]byte, error) {
	if b.MerkleRoot == nil {
		return nil, errors.New("merkle root not set")
	}

	header := struct {
		Index        int
		Timestamp    int64
		PreviousHash []byte
		MerkleRoot   []byte
	}{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		PreviousHash: b.PreviousHash,
		MerkleRoot:   b.MerkleRoot,
	}

	bytes, err := json.Marshal(header)
	if err != nil {
		return nil, err
	}

	return crypto.Hash(bytes), nil
}

func (b *Block) Finalize(node *identity.NodeIdentity) error {
	if len(b.Transactions) == 0 {
		return errors.New("block must contain transactions")
	}

	var txHashes [][]byte

	for _, tx := range b.Transactions {
		hash, err := tx.Hash()
		if err != nil {
			return err
		}
		txHashes = append(txHashes, hash)
	}

	b.MerkleRoot = ComputeMerkleRoot(txHashes)

	hash, err := b.computeBlockHash()
	if err != nil {
		return err
	}

	b.Hash = hash
	b.Validator = node.NodeID

	signature, err := node.Sign(hash)
	if err != nil {
		return err
	}

	b.Signature = signature
	return nil
}

func (b *Block) Verify(signer crypto.Signer, validatorPublicKey []byte) (bool, error) {

	if b.Hash == nil {
		return false, errors.New("block hash missing")
	}

	// 1️⃣ Verify every transaction signature
	for _, tx := range b.Transactions {
		validTx, err := tx.Verify(signer)
		if err != nil || !validTx {
			return false, nil
		}
	}

	// 2️⃣ Recompute Merkle root
	var txHashes [][]byte
	for _, tx := range b.Transactions {
		hash, err := tx.Hash()
		if err != nil {
			return false, err
		}
		txHashes = append(txHashes, hash)
	}

	recomputedMerkle := ComputeMerkleRoot(txHashes)

	if string(recomputedMerkle) != string(b.MerkleRoot) {
		return false, nil
	}

	// 3️⃣ Recompute block header hash
	expectedHash, err := b.computeBlockHash()
	if err != nil {
		return false, err
	}

	if string(expectedHash) != string(b.Hash) {
		return false, nil
	}

	// 4️⃣ Verify block signature
	return signer.Verify(validatorPublicKey, b.Hash, b.Signature), nil
}