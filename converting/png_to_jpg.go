package converting

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"png_to_jpg/utils"
	"strings"

	"github.com/sunshineplan/imgconv"
)

func getPngPaths(path string) []string {
	result := []string{}
	allowedExtension := "png"
	isDir, err := utils.CheckIsDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if isDir {
		utils.AppendPathsFromDir(&result, path, allowedExtension)
	} else {
		extension, _ := utils.GetExtenstion(path)
		if extension == allowedExtension {
			result = append(result, path)
		}
	}

	return result
}

func getNewFilePath(sourceDirPath string, resultDirPath string, path string) string {
	return strings.Replace(path, sourceDirPath, resultDirPath, 1)
}

func PngToJpg(sourceDirPath string, resultDirPath string) {
	pngPaths := getPngPaths(sourceDirPath)
	err := os.MkdirAll(resultDirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create resultDirPath: %v", err)
	}

	fmt.Println("png files: ", pngPaths)

	for _, pngPath := range pngPaths {
		src, err := imgconv.Open(pngPath)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}

		_, others := utils.GetExtenstion(pngPath)
		newFileOthers := getNewFilePath(sourceDirPath, resultDirPath, others)
		newFilePath := fmt.Sprintf("%s.jpg", newFileOthers)

		newFile, err := os.OpenFile(
			newFilePath,
			os.O_CREATE|os.O_RDWR|os.O_TRUNC,
			os.FileMode(0644),
		)
		if err != nil {
			log.Fatalf("failed to create newFile: %v", err)
		}

		w := bufio.NewWriter(newFile)
		imgconv.Write(w, src, &imgconv.FormatOption{Format: imgconv.JPEG})
		newFile.Close()

		fmt.Println("png to jpg file: ", newFilePath)
	}
}
