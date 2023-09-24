package models

type BlockChain struct {
	Blocks []Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []Block{
			*NewGenesisBlock(),
		},
	}
}

func (bc *BlockChain) WriteBlock(b *Block) {
	bc.Blocks = append(bc.Blocks, *b)
}

func (bc *BlockChain) GetBlockByBookId(id string) (int, *Block) {
	for index, block := range bc.Blocks {
		if block.Data.IsGenesis {
			continue
		}
		if block.Data.Book.ID == id {
			return index, &block
		}
	}
	return -1, nil
}
func (bc *BlockChain) DeleteBlockAtIndex(index int) bool {
	if len(bc.Blocks) > 1 && index > 0 && index < len(bc.Blocks) {
		bc.Blocks = append(bc.Blocks[0:index], bc.Blocks[index+1:]...)
		return true
	}
	return false
}
