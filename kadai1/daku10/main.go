package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.dom/daku10/gopherdojo-studyroom/kadai1/daku10/converter"
)

func main() {
	target := flag.String("target", "", "target directory")
	in := flag.String("in", ".jpg", "input file type(should not need prefix dot...")
	out := flag.String("out", ".png", "output file type")
	flag.Parse()
	t := filepath.Clean(*target)
	file, err := os.Stat(t)
	if err != nil {
		panic(err)
	}
	if !file.IsDir() {
		panic("should be directory")
	}	
	converter.Convert(t, *in, *out)
}
