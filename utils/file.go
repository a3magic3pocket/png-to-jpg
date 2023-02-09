package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func CheckIsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func RemoveLastSlice(path string) string {
	re := regexp.MustCompile(`\/{2,}`)
	removed := re.ReplaceAllString(path, "/")

	if removed[len(removed)-1:] == "/" {
		return removed[:len(removed)-1]
	}

	return removed
}

func GetExtenstion(path string) (string, string) {
	splitted := strings.Split(path, ".")
	extenstion := splitted[len(splitted)-1]
	others := strings.ReplaceAll(path, fmt.Sprintf(".%s", extenstion), "")

	return strings.ToLower(extenstion), others
}

func AppendPathsFromDir(result *[]string, path string, allowedExtension string) {
	isDir, err := CheckIsDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if !isDir {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	parentDirPath := RemoveLastSlice(path)
	for _, file := range files {
		extension, _ := GetExtenstion(file.Name())
		if extension == allowedExtension {
			*result = append(*result, fmt.Sprintf("%s/%s", parentDirPath, file.Name()))
		}
	}
}
