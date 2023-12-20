package main

import (
	"flag"
	"log"
	"stori/pkg/processor/application"
)

func main() {
	file := getFlags()
	log.Println(file)

	res, err := application.New(file, application.ProcessorFile).Process()

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
