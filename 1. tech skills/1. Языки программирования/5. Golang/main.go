package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Напишите свою заметку...")
	
	var note string

	fmt.Scan(&note)
	
	file, err := os.Create("file1")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString(note)
}

/*
func getTitleFromFile() {
	file, err := os.Open("Default_title.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := scanner.ReadString('\n')

	title := int(line)
}
*/
