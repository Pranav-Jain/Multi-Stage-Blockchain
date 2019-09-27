package main

import (
	"bytes"
	"crypto/sha256"
	// "fmt"
	"math"
	"math/big"
	"encoding/binary"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 20
const t = 10000

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	bc  *Blockchain
	stage  *Stage
	target int
	// b_digest []byte
	s_digest []byte
	stage_number int
	block_number int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Blockchain, s *Stage, stage_number int, block_number int) *ProofOfWork {

	pow := &ProofOfWork{}

	if stage_number == 0 {
		pow = &ProofOfWork{b, s, t, []byte{}, stage_number, block_number}
	} else {
		pow = &ProofOfWork{b, s, t, b.blocks[block_number].stages[stage_number-1].Digest, stage_number, block_number}
	}

	return pow
}


func (pow *ProofOfWork) prepareData(nonce int) []byte {
	
	data := []byte{}
	if pow.stage_number == 0 {
		data = bytes.Join(
			[][]byte{
				pow.bc.blocks[pow.stage_number].Digest,
				IntToHex(int64(pow.block_number)),
				pow.stage.Target,
				IntToHex(int64(t)),
				IntToHex(pow.stage.Timestamp),
				IntToHex(int64(pow.stage.Reward)),
				IntToHex(int64(nonce)),
			},
			[]byte{},
		)
	} else {
		data = bytes.Join(
			[][]byte{
				pow.bc.blocks[pow.stage_number].Digest,
				IntToHex(int64(pow.block_number)),
				IntToHex(int64(t)),
				IntToHex(pow.stage.Timestamp),
				IntToHex(int64(pow.stage.Reward)),
				IntToHex(int64(nonce)),
			},
			[]byte{},
		)
	}

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	// fmt.Printf("\nMining the block containing \"%s\"\n", pow.stage.Target)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		// fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if binary.BigEndian.Uint32(hash[:]) < t {
			break
		} else {
			nonce++
		}
	}
	// fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
// func (pow *ProofOfWork) Validate() bool {
// 	var hashInt big.Int

// 	data := pow.prepareData(pow.stage.Nonce)
// 	hash := sha256.Sum256(data)
// 	hashInt.SetBytes(hash[:])

// 	isValid := hashInt.Cmp(pow.target) == -1

// 	return isValid
// }
