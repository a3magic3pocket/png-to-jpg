package utils

import (
	"fmt"
	"testing"
)

func TestCheckIsDir(t *testing.T) {
	isDir, err := CheckIsDir("wrong-path")
	if err == nil {
		t.Fatal(err)
	}
	if isDir {
		t.Fatal("dir not exists but return value is true")
	}

	isDir, err = CheckIsDir("..")
	if err != nil {
		t.Fatal(err)
	}
	if !isDir {
		t.Fatal("dir exists but return value is false")
	}
}
func TestRemoveLastSlash(t *testing.T) {
	result := RemoveLastSlash("/test/")
	if result != "/test" {
		t.Fatal("can not remove last slash from input string")
	}

	result = RemoveLastSlash("/test///")
	if result != "/test" {
		t.Fatal("failed to remove multiple slashes at end of string")
	}

	result = RemoveLastSlash("/test2//test///")
	if result != "/test2/test" {
		t.Fatal("failed to remove multiple slashs at string")
	}

}
func TestGetExtenstion(t *testing.T) {
	extension, others := GetExtenstion("./my/test.file")
	if extension != "file" {
		t.Fatal("failed to extract correct extension")
	}

	fmt.Println("others", others)
	if others != "./my/test" {
		t.Fatal("failed to extract correct others")
	}

	extension, others = GetExtenstion("./my/test.CSV")
	if extension != "csv" {
		t.Fatal("failed to extract correct extension")
	}

	fmt.Println("others", others)
	if others != "./my/test" {
		t.Fatal("failed to extract correct others")
	}
}
func TestAppendPathsFromDir(t *testing.T) {
	result := []string{}
	AppendPathsFromDir(&result, ".", []string{"go"})

	isSuccessful := false

	for _, path := range result {
		if path == "./file_test.go" {
			isSuccessful = true
			break
		}
	}

	if !isSuccessful {
		t.Fatal("failed to find ./file_test.go from ..")
	}
}
