package converting

import (
	"errors"
	"log"
	"os"
	"png_to_jpg/utils"
	"testing"
)

func TestPngToJpg(t *testing.T) {
	// Init dir
	sourceDirPath := "./test_source"
	resultDirPath := "./test_result"
	for _, dirPath := range []string{sourceDirPath, resultDirPath} {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	PngToJpg(sourceDirPath, resultDirPath)

	result := []string{}
	utils.AppendPathsFromDir(&result, resultDirPath, "jpg")
	for _, resultPath := range result {
		if _, err := os.Stat(resultPath); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("failed to find result path: %v", os.ErrNotExist)
		}
	}

	// Clean
	os.RemoveAll(resultDirPath)
}
