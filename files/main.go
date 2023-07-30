package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	t, err := f.Write([]byte("Hello World"))

	if err != nil {
		panic(err)
	}

	fmt.Println(t, "bytes written successfully")

	defer f.Close()

	file, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Opened test.txt")

	reader := bufio.NewReader(file)

	buffer := make([]byte, 8)

	for {
		count, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Printf("read %d bytes: %q\n", count, buffer[:count])
	}

	os.Remove("test.txt")

}
