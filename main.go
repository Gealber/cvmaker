package main

import (
	"log"

	"github.com/Gealber/cvmaker/maker"
)

func main() {
	cvMaker := maker.NewCVMaker()

	log.Println("Creating cv...")

	if err := cvMaker.Generate(); err != nil {
		log.Fatal(err)
	}
}
