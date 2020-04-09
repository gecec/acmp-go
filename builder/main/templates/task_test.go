package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"bufio"
)

func TestSolution(t *testing.T){
	//iterate by list of files input*, answers*
	var tests = [] struct{
		input, answer string
	}{
		{"input1.txt", "answer1.txt"},
	}

	for _,tt := range tests {
		t.Run(tt.input, func(t *testing.T){
			if err := os.Remove("output.txt"); err != nil{
				log.Println(err)
			}

			Solution(tt.input, "output.txt")

			answers := readStringsFromFile(tt.answer)

			results := readStringsFromFile("output.txt")
			fmt.Println("answers = ", answers)
			fmt.Println("results = ", results)
			if results[0] != answers[0]{
				t.Errorf("got %s, want %s", results[0], answers[0])
			}
		})

	}
}

func closeFile(f *os.File) {
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func readStringsFromFile(fileName string) []string {
	arr := make([]string, 0)

	file,err := os.Open(fileName)
	defer closeFile(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		arr = append(arr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}