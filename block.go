package blockchain

type Block struct {
	timestamp  time.Time
	lastHash   string
	hash       string
	data       string
	nonce      uint32
	difficulty int
}

func getBlock(timestamp time.Time,
	lastHash string, hash string, data string, nonce uint32, difficulty int) *Block {
	return &Block{
		timestamp:  timestamp,
		lastHash:   lastHash,
		hash:       hash,
		data:       data,
		nonce:      nonce,
		difficulty: difficulty,
	}
}

func (b *Block) toString() string {
	s := []string{" Block - ",
		"\nTimestamp : " + b.timestamp.String(),
		"\nLast Hash :" + b.lastHash,
		"\nHash : " + b.hash,
		"\nNonce: " + fmt.Sprint(b.nonce),
		"\nDifficulty: " + strconv.Itoa(b.difficulty),
		"\nData : " + b.data}
	var x = strings.Join(s, " ")
	fmt.Println(x)
	return x
}

func (b *Block) test() string {
	return "hi"
}

func (b *Block) genesisBlock() Block {
	return Block{
		timestamp:  time.Now(),
		lastHash:   "",
		hash:       "",
		data:       "",
		nonce:      0,
		difficulty: 0,
	}
}