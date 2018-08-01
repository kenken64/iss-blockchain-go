
### Get Golang Web middleware - GIN
* Install GIN dependencies to your project home directory, further feature please refer to GIN documention (https://github.com/gin-gonic/gin)

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

alternatively 

```
go run blockchain.go
```

### Testing

Launch postman to test the following end point 

* GET http://localhost:3005/blocks
* GET http://localhost:3005/is-chain-valid
* POST http://localhost:3005/pay

Payload for the above POST

```
{
	"From": "Bala",
	"To": "Chuk",
	"Amount": 500
}
```


### Golang Notes

* short variable declaration := 
* Appending to a slice
⋅⋅⋅ func append(s []T, vs ...T) []T
* Go has pointers. A pointer holds the memory address of a value.

⋅⋅⋅The type *T is a pointer to a T value. Its zero value is nil.

⋅⋅⋅var p *int
⋅⋅⋅The & operator generates a pointer to its operand.

⋅⋅⋅i := 42
⋅⋅⋅p = &i
⋅⋅⋅The * operator denotes the pointer's underlying value.

⋅⋅⋅fmt.Println(*p) // read i through the pointer p
⋅⋅⋅*p = 21         // set i through the pointer p
⋅⋅⋅This is known as "dereferencing" or "indirecting".

⋅⋅⋅Unlike C, Go has no pointer arithmetic.

* sprintf is string formatting

* Format use for time.Time parsing RFC3339     = "2006-01-02T15:04:05Z07:00"
