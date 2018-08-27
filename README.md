### Workshop Google Codelabs
[Workshop](https://bit.ly/2MHxKmp)

### Pre-requisite
* Install NodeJS for Gomon
* Install Golang (https://golang.org/dl/)
* Install Visual Studio Code as IDE (https://code.visualstudio.com/)

### Get Golang Web middleware - GIN
* Install GIN library to your project home directory, further feature please refer to GIN documentation (https://github.com/gin-gonic/gin)

```
go get github.com/gin-gonic/gin
```

### Get Golang Websocket - GORILLA
* Install GORILLA libary to your project home directory, further feature please refer to GORILLA documentation (http://www.gorillatoolkit.org/pkg/websocket)

```
go get github.com/gorilla/websocket
```

### Get ripeMD160 
* Install package ripemd160 implements the RIPEMD-160 hash algorithm.

```bash
go get golang.org/x/crypto/ripemd160
```

### Get fast key value database
* Install bolt package for fast key value database

```bash
go get github.com/boltdb/bolt
```

Error will be thrown if the package is not installed
```
blockchain.go:17:2: cannot find package "github.com/boltdb/bolt" in any of:
	/usr/local/Cellar/go/1.10.3/libexec/src/github.com/boltdb/bolt (from $GOROOT)
	/Users/phangty/code/go/src/github.com/boltdb/bolt (from $GOPATH)
```

### Utilities required
install gomon (https://github.com/johannesboyne/gomon)

```
npm install -g go-mon
```

### How to start the blockchain app

Go to your terminal and execute the following either one of the command line

```
gomon blockchain.go
```

alternatively 

```
go run blockchain.go
```

### Compile the golang source code into binary file, execute the program in binary mode
```
go build blockchain.go utils.go base58.go
chmod +x blockchain
./blockchain -h localhost:3001 -db node1.db
./blockchain -h localhost:3002 -d localhost:3001
```

### Node synchornization using client server rather p2p

Running master node with bolt db argument
```bash
go run blockchain.go utils.go base58.go -h localhost:3001 -db node1.db
```

Running child node with bolt db argument
```bash
go run blockchain.go utils.go base58.go -h localhost:3002 -d localhost:3001 -db node2.db
```


### Get the go lang dump from the binary file
```bash
go tool objdump -S blockchain
```

### Testing Endpoint

Launch postman to test the following end point 

* GET http://localhost:3001/blocks
* GET http://localhost:3001/wallets
* GET http://localhost:3001/new-wallet?Balance=1000
* GET http://localhost:3001/ws
* GET http://localhost:3001/is-chain-valid
* POST http://localhost:3001/pay

Payload for the above POST

```
{
	"From": "Bala",
	"To": "Chuk",
	"Amount": 500
}
```

### Generate Google Codelabs tool
```bash
claat export -prefix "" blockchain-iss.md
```
