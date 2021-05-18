package cipher

import (
    "crypto/rand"
    "math/big"
)

// ************************************************************************************************
// Util
// ************************************************************************************************

const padchars = "abcdefghijklmnopqrstuvwxyz"

func RandInt(max int) int {
	if n, err := rand.Int(rand.Reader, big.NewInt(int64(max))); err != nil {
		panic("Error generating random integer.")
	} else {
		return int(n.Int64())
	}
}

func PadChar() byte {
	return padchars[RandInt(len(padchars))]
}

// ************************************************************************************************
// Cipher
// ************************************************************************************************

func RailFenceCipherEncrypt(plaintext string) string {
	padding := 4 - (len(plaintext) % 4)
    if padding == 4 {
        padding = 0
    }
	plainb := []byte(plaintext)
	for i := 0; i < padding; i++ {
		plainb = append(plainb, PadChar())
	}
	ciphertext := make([]byte, len(plainb))
	j := 0
	for i := 0; i < len(plainb); i = i + 2 {
		ciphertext[j] = plainb[i]
		j++
	}
	for i := 1; i < len(plainb); i = i + 2 {
		ciphertext[j] = plainb[i]
		j++
	}
	return string(ciphertext)
}

func RailFenceCipherDecrypt(ciphertext string) string {
	jump := len(ciphertext) / 2
	plaintext := make([]byte, len(ciphertext))
	j := 0
	for i := 0; i < jump; i++ {
		plaintext[j] = ciphertext[i]
		j++
		plaintext[j] = ciphertext[i+jump]
		j++
	}
	return string(plaintext)
}
