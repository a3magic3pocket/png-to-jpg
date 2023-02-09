# png-to-jpg
Convert png file to jpg.
if jpg exists, jpg will be smaller.

## How to run
- ```bash
    # Build
	go build -o png_to_jpg

	# and Run build file
    ./png_to_jpg [source directory of png files] [result directory of converted jpg files]

    # or Run immediately
	go run main.go [source directory of png files] [result directory of converted jpg files]
    ```

## Install dependencies
- ```bash
    go mod tidy
    ```

## Test
- ```bash
    go test ./...
    ```