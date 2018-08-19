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
go build blockchain.go
chmod +x blockchain
./blockchain -h localhost:3001
./blockchain -h localhost:3002 -d localhost:3001
```

### Node synchornization using client server rather p2p
```
go run blockchain.go -h localhost:3001
go run blockchain.go -h localhost:3002 -d localhost:3001
```


### Testing Endpoint

Launch postman to test the following end point 

* GET http://localhost:3001/blocks
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
