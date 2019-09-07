package main

import (
	"time"
)

type Stage struct {
	Target    []byte
	Timestamp int64
	Address   []byte
	Reward    int
	Digest    []byte
	Nonce     int
}

// NewBlock creates and returns Block
func NewStage(data string, prevStageHash []byte) *Stage {
	// block := &Block{[]byte(data), prevBlockHash, []byte{}, number, []struct{}}
	stage := &Stage{[]byte(data), time.Now().Unix(), []byte(data), 1, []byte{}, 0}
	pow := NewProofOfWork(stage)
	nonce, hash := pow.Run()

	stage.Digest = hash[:]
	stage.Nonce = nonce

	return stage
}