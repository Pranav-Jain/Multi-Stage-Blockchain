package main

import (
	"encoding/binary"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Number        int
	Digest        []byte
	stages		  []*Stage
}


// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddStage(data string, stage_no int) {
	prevStage := bc.blocks[len(bc.blocks) - k + stage_no]
	newStage := NewStage(data, prevStage.Digest, stage_no)
	bc.blocks[len(bc.blocks)-1].stages = append(bc.blocks[len(bc.blocks)-1].stages, newStage)

	if stage_no == k-1 {
		bc.blocks[len(bc.blocks)-1].Digest = bc.blocks[len(bc.blocks)-1].stages[k-1].Digest
	}
}


func (bc *Blockchain) NewBlock(number int) *Block {
	block := &Block{number, []byte{}, []*Stage{}}
	bc.blocks = append(bc.blocks, block)
	return block
}


func NewBlockchain() *Blockchain {
	bc := &Blockchain{[]*Block{}}
	for i := 0; i < k; i++ {
		block := &Block{i, []byte{}, []*Stage{}}
		digest := make([]byte, 8)
		binary.LittleEndian.PutUint64(digest, uint64(0))
		if i !=0 {
			prevBlock := bc.blocks[len(bc.blocks) - i]
			digest = prevBlock.Digest
		}
		bc.blocks = append(bc.blocks, block)
		newStage := NewGenesis(bc, "Genesis Block", digest, i)
		block.stages = append(block.stages, newStage)
	}	
	return bc
}

