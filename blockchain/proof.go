package blockchain

import (
	"math/big"
)

const Difficulty = 10;

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork  {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return &ProofOfWork{block, target}
}