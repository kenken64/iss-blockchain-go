package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Block struct {
	Index      int         `json:"index"`
	Timestamp  time.Time   `json:"date"`
	LastHash   string      `json:"prevHash"`
	Hash       string      `json:"hash"`
	Data       interface{} `json:"transaction"`
	nonce      uint32
	difficulty int
}

type Blocks []*Block

type Chain interface {
	AddBlock(block *Block)
	GetLastBlock() *Block
	GetBlocks() Blocks
	IsChainValid() bool
}

type BlockChain struct {
	Blocks Blocks
}

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Date   time.Time
}

func NewBlock(index int, data interface{},
	timestamp time.Time) *Block {
	return &Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
	}
}

func NewTransaction(from string, to string, amount float64) *Transaction {
	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
		Date:   time.Now(),
	}
}

func (b *Block) createHash() string {
	d := fmt.Sprintf("%v%v%v%v", b.Index, b.LastHash, b.Timestamp, b.Data)
	h := sha256.New()
	h.Write([]byte(d))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func NewBlockChain() Chain {
	date, error := time.Parse(time.RFC3339, "2018-08-20T22:08:41+00:00")
	if error != nil {
		panic(error)
	}
	// create genesis block
	genesisBlock := NewBlock(0, "Genesis Block", date)
	genesisBlock.LastHash = "0"
	genesisBlock.Hash = genesisBlock.createHash()

	var blocks Blocks
	blocks = append(blocks, genesisBlock)

	return &BlockChain{
		Blocks: blocks,
	}
}

func (b *BlockChain) AddBlock(block *Block) {
	block.LastHash = b.GetLastBlock().Hash
	block.Hash = block.createHash()
	b.Blocks = append(b.Blocks, block)
}

func (b *BlockChain) GetLastBlock() *Block {
	return b.Blocks[len(b.Blocks)-1]
}

func (b *BlockChain) GetBlocks() Blocks {
	return b.Blocks
}

func (b *BlockChain) IsChainValid() bool {
	for i := 1; i < len(b.Blocks); i++ {
		currentBlock := b.Blocks[i]
		prevBlock := b.Blocks[i-1]

		if currentBlock.Hash != currentBlock.createHash() {
			return false
		}

		if currentBlock.LastHash != prevBlock.Hash {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("[ Starting Blockchain API Server ]")
	// Instantiate a new blockchain
	blockchain := NewBlockChain()
	r := gin.Default()

	r.GET("/blocks", func(c *gin.Context) {
		currentBlocks := blockchain.GetBlocks()
		var interfaceSlice []interface{} = make([]interface{}, len(currentBlocks))
		for i := range currentBlocks {
			fmt.Println(currentBlocks[i].Index)
			fmt.Println(currentBlocks[i].Timestamp)
			fmt.Println(currentBlocks[i].Hash)
			fmt.Println(currentBlocks[i].LastHash)
			fmt.Println(currentBlocks[i].Data)
			blockJson := &Block{
				Index:     currentBlocks[i].Index,
				Timestamp: currentBlocks[i].Timestamp,
				LastHash:  currentBlocks[i].LastHash,
				Hash:      currentBlocks[i].Hash,
				Data:      currentBlocks[i].Data,
			}
			interfaceSlice[i] = blockJson
		}
		c.JSON(http.StatusOK, gin.H{"blocks": interfaceSlice})
	})

	r.POST("/pay", func(c *gin.Context) {
		var incomingTransaction Transaction
		c.BindJSON(&incomingTransaction)
		transfer := incomingTransaction
		lengthOfChain := len(blockchain.GetBlocks())
		indexblock := lengthOfChain + 1
		transferBlock := NewBlock(indexblock, transfer, time.Now())
		blockchain.AddBlock(transferBlock)
		c.JSON(http.StatusOK, gin.H{"transferedValidity": blockchain.IsChainValid()})
	})

	r.GET("/is-chain-valid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"chain-validity": blockchain.IsChainValid()})
	})
	hostNamePort := os.Getenv("BLOCKCHAIN_API_HOSTNAMEPORT")
	fmt.Println("Hostname and Port ", hostNamePort)
	if hostNamePort == "" {
		hostNamePort = "localhost:3005"
		fmt.Println("default Hostname and Port ", hostNamePort)
	}
	
	r.Run(hostNamePort)

}
