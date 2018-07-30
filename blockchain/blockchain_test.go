package blockchain_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/kenken64/iss-blockchain-go/blockchain"
)

func TestGenesisBlock(t *testing.T) {
	fmt.Print("test > ")

	b := NewBlock(time.Now(), "", "", "", 0, 0)
	fmt.Println(b.toString())
}
