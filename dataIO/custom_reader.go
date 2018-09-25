package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

type asciiReader struct {
	src []byte
	cur int
	err error
}

// newASCIIReader is used to take under control the src field when it is empty. A new asciiReader should be created by using this method.
func newASCIIReader(src string) *asciiReader {
	if len(src) == 0 {
		return &asciiReader{
			src: make([]byte, 0),
			err: io.EOF,
		}
	}

	return &asciiReader{
		src: []byte(src),
	}
}

func (a *asciiReader) Read(p []byte) (int, error) {
	var bytesASCII int

	if len(p) == 0 {
		return 0, nil
	}
	if a.err != nil {
		return 0, a.err
	}

	for bytesASCII < len(p) {
		word, sizeWord := utf8.DecodeRune(a.src[a.cur:])

		if word == utf8.RuneError {
			a.err = errors.New("An error was gotten trying to report an Unicode character")
			return bytesASCII, nil
		}

		a.cur += sizeWord

		if sizeWord == 1 {
			p[bytesASCII] = byte(word)
			bytesASCII += sizeWord
		}

		if len(a.src[a.cur:]) == 0 {
			a.err = io.EOF
			return bytesASCII, nil
		}
	}
	return bytesASCII, nil
}

func main() {
	input := newASCIIReader("It is an example in español")
	buf := make([]byte, 10)
	for {
		size, err := input.Read(buf)
		if err != nil {
			if size > 0 {
				fmt.Println(string(buf[:size]))
			}
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(buf[:size]))
	}
}
