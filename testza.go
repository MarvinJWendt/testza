package testza

import (
	"math/rand"
	"time"
)

var randomSeed int64

func init() {
	randomSeed = time.Now().UnixNano()
	rand.Seed(randomSeed)
}
