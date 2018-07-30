package blockchain

import (
	"time"
)
type Block struct {
	timestamp time.Time
	lastHash string
	hash string
	nonce string
	difficulty int
}