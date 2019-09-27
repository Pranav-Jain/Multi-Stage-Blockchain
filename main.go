package main

import (
	"fmt"
	"time"
	// "strconv"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const k = 5
const general_blocks = 10

var pts = make(plotter.XYs, 100)
var count = 0
var t1 = time.Now()

func pow(bc *Blockchain, stage *Stage, j int, i int) {
	pow := NewProofOfWork(bc, stage, j, i)
	nonce, hash := pow.Run()

	stage.Digest = hash[:]
	stage.Nonce = nonce

	// t2 := time.Now()
	pts[j*(i-k) + j].X = time.Now().Sub(t1).Seconds()
	pts[j*(i-k) + j].Y = float64(j*(i-k) + j)
	count += 1
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

	// t1 = time.Now()
	// sequential(bc)
	// pts_seq := pts

	// pts1 := make(plotter.XYs, 100)
	// pts = pts1
	// count = 0

	t1 = time.Now()
	pipelined(bc, 0, k)
	pts_pipe := pts

	
	// t4 := time.Now()
	// fmt.Printf("Sequential took seconds: %f\n", t4.Sub(t3).Seconds())



	// randomPoints returns some random x, y points
	// with some interesting kind of trend.

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Time Taken"
	p.X.Label.Text = "Time Elapsed"
	p.Y.Label.Text = "Number of Stages"

	err = plotutil.AddScatters(p,
		"Pipelined", pts_pipe)
		// "Sequential", pts_seq)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

}
