package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	// for each letter read from io.Reader
	for i := 0; i < len(p); i++ {

		// if the letter's index is between A - N, add 13 to its index
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			p[i] += 13

			// if the letter's index is between M - Z, subtract 13 from its index
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return
}

// test to make sure it works
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
	You cracked the code!
*/
