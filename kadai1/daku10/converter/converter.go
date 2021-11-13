package converter

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

func Convert(targetDirPath string, in string, out string) {
	err := filepath.Walk(targetDirPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == in {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				img, _, err := image.Decode(file)
				if err != nil {
					return err
				}
				newFile, err := os.Create(path + out)
				if err != nil {
					return err
				}
				png.Encode(newFile, img)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
