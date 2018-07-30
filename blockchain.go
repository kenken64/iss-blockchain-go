package blockchain

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

package blockchain

type Blockchain struct {
	chain	[]*Block `json:"chain"`
}