package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Input:
// - task number
// - templates/task.go template
// - templates/task_test.go template
//
// ToDo:
// builder <task number>
// - create folder t<task number>
// - create folder t<task number>/main
// - copy template files to t<task number>/main
// create files input1.txt, answer1.txt
func main (){
	taskNum := os.Args[1]
	_, err := strconv.Atoi(taskNum)
	if err != nil {
		log.Fatal("Wrong argument [%s]", taskNum, err)
		os.Exit(1)
	}

	curDir, _ := os.Getwd()

	os.Chdir("../..")
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rootDir)
	taskPath := filepath.Join(rootDir, "t" + taskNum, "main")

	err = os.MkdirAll(taskPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	copyFile(
		filepath.Join(curDir, "templates", "task.go"),
		filepath.Join(taskPath, "task.go"))

	copyFile(
		filepath.Join(curDir, "templates", "task_test.go"),
		filepath.Join(taskPath, "task_test.go"))


	input := createFile(filepath.Join(taskPath, "input1.txt"))
	defer input.Close()

	answer := createFile(filepath.Join(taskPath, "answer1.txt"))
	defer answer.Close()
}

func copyFile(sourceFilePath, destFilePath string){
	log.Println("source = ", sourceFilePath)
	log.Println("target = ", destFilePath)

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer sourceFile.Close()

	destFile := createFile(destFilePath)
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	destFile.Sync()
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return file
}