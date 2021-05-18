package main

import (
	"fmt"
	"os"
)

// ************************************************************************************************
// main()
// ************************************************************************************************

func main() {
	terminal := NewTerminal(os.Stdin, os.Stdout)
	for {
		terminal.Prompt()
		line := terminal.ReadLine()
		if len(line) == 0 {
			break
		}
		plaintext := line[:len(line)-1] // remove terminating newline
		ciphertext := RailFenceCipherEncrypt(plaintext)
		fmt.Printf("%s => %s => %s\n", plaintext, ciphertext, RailFenceCipherDecrypt(ciphertext))
	}
}
