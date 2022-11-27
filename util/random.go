package util

import (
	"math/rand"
	"time"
)

func Get_Random_Int(start, end int) int {
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	randNum := rand.Intn(end-start) + start
	return randNum
}
