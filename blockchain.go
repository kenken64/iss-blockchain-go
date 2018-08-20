package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ripemd160"
)

// data structure of the block
type Block struct {
	Index      int         `json:"index"`
	Timestamp  time.Time   `json:"date"`
	LastHash   string      `json:"prevHash"`
	Hash       string      `json:"hash"`
	Data       interface{} `json:"transaction"`
	nonce      uint32
	difficulty int
}

// list of blockchains
type Blocks []*Block

var mutex = &sync.Mutex{}

// interface consist all the function of the blockchain struct
type Chain interface {
	AddBlock(block *Block)
	GetLastBlock() *Block
	GetBlocks() Blocks
	IsChainValid() bool
}

type BlockChain struct {
	Blocks Blocks `json:"Blocks"`
}

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Date   time.Time
}

const version = byte(0x00)
const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
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

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base58Encode encodes a byte array to Base58
func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	// https://en.bitcoin.it/wiki/Base58Check_encoding#Version_bytes
	if input[0] == 0x00 {
		result = append(result, b58Alphabet[0])
	}

	ReverseBytes(result)

	return result
}

// Base58Decode decodes Base58-encoded data
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)

	for _, b := range input {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()

	if input[0] == b58Alphabet[0] {
		decoded = append([]byte{0x00}, decoded...)
	}

	return decoded
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// GetAddress returns wallet address
func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wshandler(w http.ResponseWriter, r *http.Request, b Chain) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	//defer conn.Close()

	go func() {
		for {
			_, message, error := conn.ReadMessage()
			if error != nil {
				fmt.Println("read: ", error)
				return
			}
			time.Sleep(2 * time.Second)
			fmt.Println("message:", string(message))
			for _, bb := range b.GetBlocks() {
				b, err := json.Marshal(bb)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteJSON(string(b))
			}

		}
	}()
}

func main() {
	fmt.Println("[ Starting Blockchain API Server ]")
	// Instantiate a new blockchain
	blockchain := NewBlockChain()
	r := gin.Default()

	help := flag.Bool("help", false, "Display Help")
	host := flag.String("h", "", "Host Address and Port")
	dest := flag.String("d", "", "Dest MultiAddr String")
	flag.Parse()
	if *help {
		fmt.Printf("This program demonstrates a simple blockchain\n\n")
		fmt.Printf("Usage: Run './blockchain -h <hostname:port> -d <peershost:port>\n")

		os.Exit(0)
	}

	blocksSize := len(blockchain.GetBlocks())
	fmt.Println("Total block size : ", blocksSize)
	fmt.Println("readsssss:", *dest)
	if *dest != "" {
		URL := "ws://" + *dest + "/ws"

		var dialer *websocket.Dialer

		conn, _, err := dialer.Dial(URL, nil)
		//defer conn.Close()
		go func() {

			if err != nil {
				fmt.Println(err)
				return
			}
			for {
				b, err := json.Marshal(blockchain.GetBlocks())
				if err != nil {
					fmt.Println(err)
					return
				}

				conn.WriteJSON(b)
				//i := Block{}
				var jsonStr = ""
				err2 := conn.ReadJSON(&jsonStr)

				if err2 != nil {
					fmt.Println("read:", err2)
					return
				}
				fmt.Println(jsonStr)
			}
		}()
	}

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request, blockchain)
	})

	r.GET("/blocks", func(c *gin.Context) {
		currentBlocks := blockchain.GetBlocks()
		c.JSON(http.StatusOK, gin.H{"blocks": currentBlocks})
	})

	r.POST("/pay", func(c *gin.Context) {
		var err error
		var incomingTransaction Transaction
		if err = c.BindJSON(&incomingTransaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}
		lengthOfChain := len(blockchain.GetBlocks())
		indexBlock := lengthOfChain + 1
		xferTransaction := NewTransaction(incomingTransaction.From, incomingTransaction.To, incomingTransaction.Amount)
		transferBlock := NewBlock(indexBlock, xferTransaction, time.Now())
		blockchain.AddBlock(transferBlock)
		c.JSON(http.StatusOK, gin.H{"transferValidity": blockchain.IsChainValid()})
	})

	r.GET("/is-chain-valid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"chain-validity": blockchain.IsChainValid()})
	})

	fmt.Println("Hostname and Port ", dest)
	var finalPortAssignment = "localhost:3005"
	if *host != "" {
		finalPortAssignment = *host
		fmt.Println("default Hostname and Port ", finalPortAssignment)
	}

	r.Run(finalPortAssignment)
}
