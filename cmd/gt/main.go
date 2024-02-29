package main

import (
	"fmt"
	"log"

	"github.com/rafaelmdm/gt/pkg/gt"
)

func main() {
	options, err := gt.GetOptions()
	if err != nil {
		log.Fatalf("Unable to get options: %v", err)
	}

	config, err := gt.NewConfig(options)
	if err != nil {
		log.Fatalf("Unable to get config: %v", err)
	}

	gt, err := gt.NewGotoCLI(config)
	if err != nil {
		log.Printf("Using default values due to: %v\n", err)
	}

	out, err := gt.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	err = gt.Save()
	if err != nil {
		log.Fatalf("Error saving file: %v", err)
	}

	fmt.Println(out)
}
