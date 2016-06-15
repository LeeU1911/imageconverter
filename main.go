package main

import (
	"github.com/leeu1911/imageconverter/converter"
	"fmt"
	"flag"
)

func main() {
	flagImg := flag.String("path", "C:/Images/image.png", "path to image file")
	flag.Parse()
	encodedStr := converter.Encode(*flagImg)
	fmt.Println(encodedStr)
}
