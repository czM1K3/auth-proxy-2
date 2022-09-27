package env

import (
	"math/rand"
	"time"
)

var secret []byte
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const possibleLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateSecret() {
	newSecret := make([]byte, 32)
	for i := range newSecret {
		newSecret[i] = possibleLetters[seededRand.Int63()%int64(len(possibleLetters))]
	}
	secret = newSecret
}

func GetSecret() []byte {
	return secret
}
