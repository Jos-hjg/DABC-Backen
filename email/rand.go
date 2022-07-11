package email

import (
	"log"
	"math/rand"
	"time"
)

var code6 = map[[6]int]bool{}

func GetRand() [6]int {
	rand.Seed(time.Now().UnixNano())
	mod := [6]int{}
	for i := 0; i < 6; i++ {
		mod[i] = rand.Intn(10)
	}
	log.Println(mod)
	code6[mod] = true
	return mod
}
