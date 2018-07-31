
### Get Golang Web middleware - GIN
* Install Golang dependencies, refer to GIN documention (https://github.com/gin-gonic/gin)

```
go get github.com/gin-gonic/gin
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

aleternatively 

```
go run blockchain.go
```

### Testing

Launch postman to test the following end point 

GET http://localhost:3005/blocks
GET http://localhost:3005/is-chain-valid
POST http://localhost:3005/pay

```
{
	"From": "Bala",
	"To": "Chuk",
	"Amount": 500
}
```