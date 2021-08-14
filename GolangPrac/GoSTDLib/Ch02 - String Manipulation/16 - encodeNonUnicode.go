package main

import (
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	f, err := os.OpenFile("out.text", os.O_CREATE|os.O_RDWR,
		os.ModePerm|os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//	Decode to unicode
	encoder := charmap.Windows1250.NewEncoder()
	writer := encoder.writer(f)
	io.WriteString(writer, "Gdańsk")
}
