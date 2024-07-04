package blockchain


type BlockChain struct {
	Blocks []*Block
}

func  InitBlockChain() *BlockChain {
	genesis := Genesis();
	blockChain := &BlockChain{[]*Block{genesis}}

	return blockChain;
}

func (chain *BlockChain) AddBlock(data string) {

	lastBlock := chain.Blocks[len(chain.Blocks) - 1]
	newBlock := CreateBlock(data, lastBlock);
	chain.Blocks = append(chain.Blocks, newBlock)
}