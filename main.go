package main

import (
	"fmt"
	"time"
	// "strconv"
	"gonum.org/v1/plot/plotter"
	"math/rand"
)

const k = 5
const general_blocks = 10

func pow(bc *Blockchain, stage *Stage, j int, i int) {
	pow := NewProofOfWork(bc, stage, j, i)
	nonce, hash := pow.Run()

	stage.Digest = hash[:]
	stage.Nonce = nonce
}

func pipelined(bc *Blockchain, j int, i int) {

	if j == k {
		return
	}
	if i == k + general_blocks {
		return
	}
	pow(bc, bc.blocks[i].stages[j], j, i)
	go pipelined(bc, j, i+1)
	pipelined(bc, j+1, i)

	return
}


func sequential(bc *Blockchain) {
	for i, block := range bc.blocks {
		for j, stage := range block.stages {
			if(i>k) {
				pow(bc, stage, j, i)
			}
		}
	}
}

func main() {
	bc := NewBlockchain()

	for i := 0; i < general_blocks; i++ {
		bc.NewBlock(i)
    	for j := 0; j < k; j++ {
      		bc.AddStage("Send 1 Coin to Ivan", j)
   		}
   }

    fmt.Printf("Genesis Blocks Created\n\n")

	// p, err := plot.New()
	// t1 := time.Now()
	// pipelined(bc, 0, k)
	// t2 := time.Now()
	// fmt.Printf("Pipelined took seconds: %f\n", t2.Sub(t1).Seconds())

	// t3 := time.Now()
	// sequential(bc)
	// t4 := time.Now()
	// fmt.Printf("Sequential took seconds: %f\n", t4.Sub(t3).Seconds())


	rnd := rand.New(rand.NewSource(1))

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) plotter.XYs {
	    pts := make(plotter.XYs, n)
	    for i := range pts {
	        if i == 0 {
	            pts[i].X = rnd.Float64()
	        } else {
	            pts[i].X = pts[i-1].X + rnd.Float64()
	        }
	        pts[i].Y = pts[i].X + 10*rnd.Float64()
	    }
	    return pts
	}

	n := 15
	scatterData := randomPoints(n)
	lineData := randomPoints(n)
	linePointsData := randomPoints(n)

	p, err := plot.New()
	if err != nil {
	    log.Panic(err)
	}
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

}
