package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
)

func getFilenameWithoutExt(file fs.FileInfo) string {
	return file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))]
}

func getFileExt(file fs.FileInfo) string {
	return filepath.Ext(file.Name())
}

func compressImage(dirname string, file fs.FileInfo) error {
	buffer, err := ioutil.ReadFile(filepath.Join("./files", file.Name()))
	if err != nil {
		return err
	}

	newFilename := getFilenameWithoutExt(file) + ".webp"
	oldFilename := file.Name()

	newFilepath := fmt.Sprintf(dirname+"/%s", newFilename)
	oldFilepath := fmt.Sprintf(dirname+"/%s", oldFilename)

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: 25})
	if err != nil {
		return err
	}

	err = os.Remove(oldFilepath)
	if err != nil {
		return err
	}

	err = bimg.Write(newFilepath, processed)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	files, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		err := compressImage("./files", file)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
