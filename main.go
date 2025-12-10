package main

import (
	"fmt"
	//"os"
)

const WIDTH, HEIGHT = 448, 224

func main() {
	fmt.Print("\r\nN")

	img := MakeTextImg("PLA\nMüll")
	_ = Convert(img)

	fmt.Print("\r\nD10") // 0-15, default 7

	fmt.Printf("\r\nGW0,0,%d,%d,", WIDTH / 8, HEIGHT)
	//os.Stdout.Write(data)

	fmt.Print("\r\nP1\r\n")
}
