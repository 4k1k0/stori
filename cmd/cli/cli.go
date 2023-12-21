package main

import (
	"embed"
	"flag"
	"log"

	"stori/internal/config"
	processor "stori/pkg/processor/application"
)

//go:embed all:assets/*
var assets embed.FS

func main() {
	file, email := getFlags()
	config.New(assets, file, email)

	res, err := processor.New().Process()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", res)
}

func getFlags() (string, string) {
	file := flag.String("file", "", "a file path")
	email := flag.String("email", "", "an email")
	flag.Parse()

	return *file, *email
}
