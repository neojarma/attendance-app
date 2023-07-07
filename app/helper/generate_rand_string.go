package helper

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenerateRandString(length int) string {
	keys := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

	res := make([]rune, length)
	for i := range res {
		res[i] = keys[rand.Intn(len(keys))]
	}

	return string(res)
}
