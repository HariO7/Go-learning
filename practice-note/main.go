package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice-note/note"
)

func main() {

	title := getUserInput("Enter the title of the note: ")
	content := getUserInput("Enter the content of the note: ")

	userNote, err := note.New(title, content)

	if err != nil {
		println(err)
		return
	}

	userNote.Display()
	userNote.Save()
}

func getUserInput(promptText string) string {
	fmt.Printf("%v ", promptText)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
