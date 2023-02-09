package converting

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"png_to_jpg/utils"
	"strings"

	"github.com/disintegration/imaging"
)

func getPngPaths(path string) []string {
	result := []string{}
	allowedExtensions := []string{"png", "jpg", "jpeg"}
	isDir, err := utils.CheckIsDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if isDir {
		utils.AppendPathsFromDir(&result, path, allowedExtensions)
	} else {
		extension, _ := utils.GetExtenstion(path)
		if utils.CheckInStringArray(extension, &allowedExtensions) {
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
		maxWidth := 800
		jpegQuality := 70

		_, others := utils.GetExtenstion(pngPath)

		// Create new File
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

		// Open origin image file
		imageFile, err := os.Open(pngPath)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}

		// Decode origin image file
		src, _, err := image.Decode(imageFile)
		if err != nil {
			log.Fatalf("failed to decode image: %v", err)
		}
		imageFile.Close()

		// Resize image if origin image is too big
		width := src.Bounds().Max.X
		if width > maxWidth {
			width = maxWidth
		}
		src = imaging.Resize(src, width, 0, imaging.Lanczos)

		// Encode resized image to jpg:
		jpegOptions := jpeg.Options{
			Quality: jpegQuality,
		}
		err = jpeg.Encode(newFile, src, &jpegOptions)
		if err != nil {
			log.Fatalf("failed to ecode image to jpg: %v", err)
		}

		newFile.Close()

		fmt.Println("png to jpg file: ", newFilePath)
	}
}
