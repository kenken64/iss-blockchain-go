package blockchain

import (
	"fmt"
	"testing"
	"time"
)

func TestGenesisBlock(t *testing.T) {
	fmt.Print("test > ")

	b := getBlock(time.Now(), "", "", "", 0, 0)
	fmt.Println(b.toString())
}
