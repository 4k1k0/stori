package main

import (
	"flag"
	"log"

	processor "stori/pkg/processor/application"
)

func main() {
	file, email := getFlags()
	res, err := processor.New(file, email).Process()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}

func getFlags() (string, string) {
	file := flag.String("file", "", "a file path")
	email := flag.String("email", "", "an email")
	flag.Parse()

	return *file, *email
}
