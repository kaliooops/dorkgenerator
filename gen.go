package main

import (
	"bufio"
	"log"
	"os"
)

//read from keywords.txt
func read_words() []string {
	file, err := os.Open("keywords.txt")

	if err != nil {
		log.Fatal("no words")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	file.Close()

	return words

}

func read_formats() []string {
	file, err := os.Open("formats.txt")

	if err != nil {
		log.Fatal("no formats")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var formats []string

	for scanner.Scan() {
		formats = append(formats, scanner.Text())
	}

	file.Close()

	return formats

}

func read_types() []string {
	file, err := os.Open("types.txt")

	if err != nil {
		log.Fatal("no types")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var types []string

	for scanner.Scan() {
		types = append(types, scanner.Text())
	}

	file.Close()

	return types

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

func main() {
	words := read_words()
	formats := read_formats()
	types := read_types()
	for _, word := range words {
		for _, format := range formats {
			for _, type_ := range types {
				WriteFile("dorks.txt", word+format+type_+"\n")
			}
		}
	}
}
