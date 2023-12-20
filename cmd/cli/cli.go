package main

import (
	"flag"
	"log"

	processor "stori/pkg/processor/application"
)

func main() {
	file := getFlags()
	log.Println(file)

	res, err := processor.New(file).Process()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}

func getFlags() string {
	file := flag.String("file", "", "a file path")
	flag.Parse()

	return *file
}
