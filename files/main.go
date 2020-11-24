package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := os.Args[1]
	open(file)
}

func open(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, fileErr := file.Read(data)
	if fileErr != nil {
		log.Fatal(err)
	}
	fmt.Printf("File Size: %v bytes\n%v", count, string(data[:count]))
}