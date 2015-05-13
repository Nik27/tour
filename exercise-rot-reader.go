package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, error := rot13.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] -= 'a'
			b[i] += 13
			b[i] %= 26
			b[i] += 'a'
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] -= 'A'
			b[i] += 13
			b[i] %= 26
			b[i] += 'A'
		}
	}
	return n, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
