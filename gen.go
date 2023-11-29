package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	words := read_file("words.txt")
	formats := read_file("formats.txt")
	types := read_file("types.txt")
	for _, word := range words {
		for _, format := range formats {
			for _, type_ := range types {
				WriteFile("dorks.txt", word+format+type_+"\n")
			}
		}
	}
}

func read_file(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("eroare")
	}

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}


func WriteFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString(content); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}
