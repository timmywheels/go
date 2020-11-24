package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct{}

func main() {
	res, err := http.Get("https://timwheeler.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	log := logger{}

	//bs := make([]byte, 99999)
	//res.Body.Read(bs)
	//fmt.Println(string(bs))
	//
	//res.Write()

	io.Copy(log, res.Body)


}

func (logger) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Bytes: %v", len(bs))
	return len(bs), nil
}