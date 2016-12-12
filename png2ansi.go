package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

const (
	SETBGCOLOR string = "\x1b[48;2"
	CRESET     string = "\033[0m"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("png2ansi usage: png2ansi <pngfile>")
		os.Exit(0)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, err)
	}
	bound := img.Bounds()

	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				fmt.Printf("%v ", CRESET)
			} else {
				fmt.Printf("%v;%v;%v;%vm ", SETBGCOLOR, uint8(r), uint8(g), uint8(b))
			}
		}
		fmt.Println("")
	}
}
