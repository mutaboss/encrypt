package main

import (
    "bufio"
    "io"
)

// ************************************************************************************************
// Terminal: Manage input and output.
// ************************************************************************************************

type Terminal struct {
	in  *bufio.Reader
	out *bufio.Writer
}

func NewTerminal(in io.Reader, out io.Writer) *Terminal {
	return &Terminal{bufio.NewReader(in), bufio.NewWriter(out)}
}

func (term *Terminal) ReadLine() string {
	if line, err := term.in.ReadString('\n'); err == io.EOF {
		if len(line) > 0 {
			return line
		} else {
			return ""
		}
	} else if err != nil {
		panic("Error reading from stdin.\n")
	} else {
		return line
	}
}

func (term *Terminal) WriteLine(str string) {
	term.out.Write([]byte(str))
}

func (term *Terminal) Prompt() {
	term.out.Write([]byte("> "))
	term.out.Flush()
}

