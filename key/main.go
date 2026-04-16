package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

// GenerateRandomKey 密钥
func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func main() {
	KeyLength := 32
	randomKey, err := GenerateRandomKey(KeyLength)
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", randomKey)
}
