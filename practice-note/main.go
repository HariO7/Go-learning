package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice-note/note"
	"example.com/practice-note/todo"
)

type Outputable interface {
	Save() error
	Display()
}

func main() {

	title := getUserInput("Enter the title of the note: ")
	content := getUserInput("Enter the content of the note: ")
	text := getUserInput("Enter your todo task")

	userNote, err := note.New(title, content)

	if err != nil {
		println(err)
		return
	}

	saveData(userNote)

	userTodo, err := todo.New(text)
	if err != nil {
		println(err)
		return
	}

	saveData(userTodo)
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

func saveData(data Outputable) error {

	err := data.Save()
	data.Display()

	if err != nil {
		println(err)
		return err
	}

	return nil
}
