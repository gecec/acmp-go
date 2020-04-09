package main

import (
	"fmt"
	"log"
	"os"
)

const (
	DEFAULT_INPUT_FILE string = "input.txt"
	DEFAULT_OUTPUT_FILE string = "output.txt"
)

func main(){
	Solution(DEFAULT_INPUT_FILE, DEFAULT_OUTPUT_FILE)
}

func Solution(inputFile, outputFile string){
	in, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer out.Close()

	var a,b int
	fmt.Fscanf(in, "%d %d", &a, &b)
	fmt.Fprintf(out, "%d", a+b)
}