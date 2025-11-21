package main

import (
	"log"

	"github.com/rojack96/jinres/gen"
)

func main() {
	if err := gen.GenerateStructsFile("status_structs_gen.go"); err != nil {
		log.Fatal(err)
	}
	log.Println("generated status_structs_gen.go")
}
