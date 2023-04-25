package main

import (
	"fmt"
	printer "mobydick/internal/format"
	"mobydick/internal/text"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите путь до файла")
		return
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("ошибка чтения файла: %v\n", err)
		return
	}

	wordFrequencies := text.ProcessText(data)
	topWords := text.FindTopNWords(wordFrequencies, 20)
	printer.PrintWords(topWords)
}
