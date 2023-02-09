# png-to-jpg
Convert png file to jpg

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