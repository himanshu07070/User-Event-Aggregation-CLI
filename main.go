package main

import (
	"flag"
	"fmt"
	"os"
	aggregate "user_events/Aggregate"
)

func main() {
	inputFile := flag.String("input", "input.json", "Input JSON file")
	outputFile := flag.String("output", "output.json", "Output JSON file")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Usage: ./aggregate_events -i input.json -o output.json")
		os.Exit(1)
	}

	aggregate.AggregateEvents(*inputFile, *outputFile)
	fmt.Println("Done!")
}
