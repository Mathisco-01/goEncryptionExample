package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)


func hash(s string) (hs string) {
	h := sha256.New()
	h.Write([]byte(s))
	hs = hex.EncodeToString(h.Sum(nil))
	return hs
}

func hashWithSalt(s string) (hs string, salt string) {

	salt_lenght := 64

	//Create random seed based on unix time
	rand.Seed(time.Now().UnixNano() * rand.Int63n(32))

	//Make saltbyte
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	saltByte := make([]byte, salt_lenght)
	for i := range saltByte {
		saltByte[i] = letters[rand.Intn(len(letters))]
	}

	//Call hash function but add salt to "s" argument
	hs = hash(s + string(saltByte))
	return hs, string(saltByte)
}