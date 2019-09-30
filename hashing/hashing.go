// Author: Mathis Van Eetvelde 			Mathisco-01

package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

type HashedLoginInfo struct {
	Hashed_username string
	Hashed_password string
	Password_salt string
}

func HashLoginInfo(username string, password string) (HLI HashedLoginInfo) {
	HLI.Hashed_username = Hash(username)
	HLI.Hashed_password, HLI.Password_salt = HashWithSalt(password)

	return HLI
}

func Hash(s string) (hs string) {
	h := sha256.New()
	h.Write([]byte(s))
	hs = hex.EncodeToString(h.Sum(nil))
	return hs
}

func HashWithSalt(s string) (hs string, salt string) {

	salt_lenght := 64

	//Create random seed based on unix time
	rand.Seed(time.Now().UnixNano() * rand.Int63n(32))

	//Make saltbyte
	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	saltByte := make([]byte, salt_lenght)
	for i := range saltByte {
		saltByte[i] = chars[rand.Intn(len(chars))]
	}

	//Call hash function but add salt to "s" argument
	hs = Hash(s + string(saltByte))
	return hs, string(saltByte)
}