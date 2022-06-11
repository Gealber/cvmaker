package main

import (
	"log"

	"github.com/Gealber/cvmaker/maker"
)

func main() {
	cvMaker, err := maker.NewCVMaker()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating cv...")

	if err := cvMaker.Generate(); err != nil {
		log.Fatal(err)
	}
}
