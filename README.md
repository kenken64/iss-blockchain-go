### Google Codelabs Workshop Style Layout

[Workshop](https://bit.ly/2MHxKmp)

### Pre-requisite

- Install NodeJS for Gomon
- Install Golang (https://golang.org/dl/)
- Install Visual Studio Code as IDE (https://code.visualstudio.com/)

### Get Golang Web middleware - GIN

- Install GIN library to your project home directory, further feature please refer to GIN documentation (https://github.com/gin-gonic/gin)

```
go get github.com/gin-gonic/gin
```

### Get Golang Websocket - GORILLA

- Install GORILLA libary to your project home directory, further feature please refer to GORILLA documentation (http://www.gorillatoolkit.org/pkg/websocket)

```
go get github.com/gorilla/websocket
```

### Get ripeMD160

- Install package ripemd160 implements the RIPEMD-160 hash algorithm.

```bash
go get golang.org/x/crypto/ripemd160
```

### Get fast key value database

- Install bolt package for fast key value database

```bash
go get github.com/boltdb/bolt
```

Error will be thrown if the package is not installed

```
blockchain.go:17:2: cannot find package "github.com/boltdb/bolt" in any of:
	/usr/local/Cellar/go/1.10.3/libexec/src/github.com/boltdb/bolt (from $GOROOT)
	/Users/phangty/code/go/src/github.com/boltdb/bolt (from $GOPATH)
```

### How to update all installed golang packages

```bash
go get -u all
```

### Utilities required

install gomon (https://github.com/johannesboyne/gomon)

```
npm install -g go-mon
```

### Compile the golang source code into binary file, execute the program in binary mode

```
go build blockchain.go utils.go base58.go
chmod +x blockchain
```

| Argument | Description                                     |     |     |     |
| -------- | ----------------------------------------------- | --- | --- | --- |
| -h       | bind to current node's hostname and port number |     |     |     |
| -d       | Connect to the master node                      |     |     |     |
| -db      | Directory of the local fast data storage        |     |     |     |

Node 1

```bash
./blockchain -h localhost:3001 -db node1.db
```

Node 2

```bash
./blockchain -h localhost:3002 -d localhost:3001 -db node2.db
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

### How to start the blockchain app

Go to your terminal and execute the following either one of the command line

```
gomon blockchain.go utils.go base58.go
```

alternatively

```
go run blockchain.go utils.go base58.go
```

### Get the go lang dump from the binary file

```bash
go tool objdump -S blockchain
```

### Testing Endpoint

1. Launch Postman select GET enter http://localhost:3001/new-wallet?Balance=1000 on the url field
2. Copy the public key address and save it somewhere on your notepad
3. Repeat step 1 and 2 again
4. Check the current blocks (aka chain) creating another test case GET http://localhost:3001/blocks
5. Perform transaction using two newly created wallet. Replace the From and To with your newly generated public key address. Input a valid from amount. POST http://localhost:3001/pay
6. If you have a child node up and running use Postman and retrieve the blocks GET http://localhost:3002/blocks. The blocks record should be in sync.

Payload for the above POST

```
{
	"From": "1255ixJZuuibAhVRMX4W8Wc3LAUv26mZ8o",
	"To": "161REKjszyKzo7Ew5hoGz1AT8wdGTJ5wj7",
	"Amount": 500
}
```

Launch postman to test the following end point

| HTTP Method | URL                                           | Description                           |     |     |
| ----------- | --------------------------------------------- | ------------------------------------- | --- | --- |
| GET         | http://localhost:3001/blocks                  | Get All Blocks                        |     |     |
| GET         | http://localhost:3001/wallets                 | Get All Wallets                       |     |     |
| GET         | http://localhost:3001/new-wallet?Balance=1000 | Create New Wallet with specify amount |     |     |
| GET         | http://localhost:3001/is-chain-valid          | Check whether the chain is valid      |     |     |
| POST        | http://localhost:3001/pay                     | Perform a pay transaction             |     |     |

### Generate Google Codelabs tool

```bash
claat export -prefix "" blockchain-iss.md
```
