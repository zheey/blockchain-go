package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash string
	Data string
	PrevHash string
}

func (b *Block) ComputeHash ()   {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computedHash := sha256.Sum256(concatenatedData)

	b.Hash = string(hex.EncodeToString(computedHash[:]))
}
