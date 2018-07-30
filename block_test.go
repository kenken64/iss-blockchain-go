package blockchain_test

import (
	"fmt"
	"testing"
	"time"

	blockchain "github.com/kenken64/iss-blockchain-go"
)

func TestGenesisBlock(t *testing.T) {
	fmt.Print("test > ")
	b := block.New(time.Now(), "", "", "", 0, 0)
	fmt.Println(b.toString())
	
	//gBlock := b.genesis()
	//fmt.Println(gBlock)
}
