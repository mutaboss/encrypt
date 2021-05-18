package main

import (
	"fmt"
	"github.com/mutaboss/encrypt/cipher"
	"github.com/mutaboss/encrypt/util"
	"os"
)

// ************************************************************************************************
// main()
// ************************************************************************************************

func main() {
	terminal := util.NewTerminal(os.Stdin, os.Stdout)
	for {
		terminal.Prompt()
		line := terminal.ReadLine()
		if len(line) == 0 {
			break
		}
		plaintext := line[:len(line)-1] // remove terminating newline
		ciphertext := cipher.RailFenceCipherEncrypt(plaintext)
		fmt.Printf("%s => %s => %s\n", plaintext, ciphertext, cipher.RailFenceCipherDecrypt(ciphertext))
	}
}
