package main

import (
	"time"
)

type Stage struct {
	number	  int
	Target    []byte
	Timestamp int64
	Address   []byte
	Reward    int
	Digest    []byte
	Nonce     int
}

// NewBlock creates and returns Block
func NewStage(data string, prevStageHash []byte, i int) *Stage {
	// block := &Block{[]byte(data), prevBlockHash, []byte{}, number, []struct{}}
	stage := &Stage{i, []byte(data), time.Now().Unix(), []byte(data), 1, []byte{}, 0}
	// pow := NewProofOfWork(stage)
	// nonce, hash := pow.Run()

	// stage.Digest = hash[:]
	// stage.Nonce = nonce

	return stage
}


func NewGenesis(bc *Blockchain, data string, prevStageHash []byte, i int) *Stage {
	// block := &Block{[]byte(data), prevBlockHash, []byte{}, number, []struct{}}
	stage := &Stage{0, []byte(data), time.Now().Unix(), []byte(data), 1, []byte{}, 0}
	pow := NewProofOfWork(bc, stage, 0, i)
	nonce, hash := pow.Run()

	stage.Digest = hash[:]
	stage.Nonce = nonce

	return stage
}