package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Split a string into chunks of the buffer length")
	buf := make([]byte, 10)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(string(buf[:n]))
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(buf[:n]))
	}
}
