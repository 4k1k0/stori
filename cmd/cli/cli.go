package main

import (
	"flag"
	"log"
)

func main() {
	file := getFlags()
	log.Println(file)
}

func getFlags() string {
	file := flag.String("file", "", "a file path")
	flag.Parse()

	return *file
}
