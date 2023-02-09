package main

import (
	"fmt"
	"os"
	"png_to_jpg/converting"
)

func help() {
	fmt.Println(`- Hwo to use?
	- build
		- go build -o png_to_jpg
	- run build file
		- ./png_to_jpg [source dir png files] [result dir converted jpg files]
	- run immediately
		- go run main.go [source dir png files] [result dir converted jpg files]
- What is png-to-jpg?
	- Convert png file to jpg file`)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		help()
		return
	}

	sourceDirPath := args[0]
	resultDirPath := args[1]
	converting.PngToJpg(sourceDirPath, resultDirPath)
}
