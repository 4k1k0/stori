package main

import (
	"flag"
	"log"
	"stori/pkg/processor/application"
	"stori/pkg/transaction"
)

func main() {
	file := getFlags()
	log.Println(file)

	res, err := application.New(file, application.ProcessorFile).Process()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	checkData()

}

func checkData() {
	foo("")
	foo("9")
	foo("1")
	foo("12.2")
	foo("+")
	foo("+1")
	foo("+1.")
	foo("+1.2")
	foo("-1.2")
	foo("-1.")
	foo("-1")
	foo("-")
	foo("-1.2s")
	foo("as-1.2")
}

func foo(d string) {
	trx := transaction.Transaction{Date: "1/11", Transaction: d}
	err := trx.Validate()
	if err != nil {
		log.Println(err)
	}
}

func getFlags() string {
	file := flag.String("file", "", "a file path")
	flag.Parse()

	return *file
}
