package blockchain_test

import (
	"fmt"
	"testing"
	"time"

	blockchain "github.com/kenken64/iss-blockchain-go"
)

func TestGenesisBlock(t *testing.T) {
	fmt.Print("test > ")

	b := blockchain.New(time.Now(), "", "", "", 0, 0)
	fmt.Println(b)
}
