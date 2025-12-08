package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("\r\nN")
	data, err := os.ReadFile("out.bin")
	if err != nil {
		panic(err)
	}

	fmt.Print("\r\nD10") // 0-15, default 7

	height := 256 // In lines (or dots)
	width := len(data) / height // In bytes


	fmt.Printf("\r\nGW0,0,%d,%d,", width, height)
	os.Stdout.Write(data)

	fmt.Print("\r\nP10\r\n")
}
