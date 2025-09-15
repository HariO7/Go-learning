package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("The title you entered is: %v \nThe content you entered is: \n\n%v\n ", note.Title, note.Content)
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	data, err := json.Marshal(note)

	if err != nil {
		fmt.Printf("Error occured while decoding:%v", err)
		return err
	}

	err = os.WriteFile(fileName, data, 0644)

	if err != nil {
		fmt.Printf("Saving file failed")
		return err
	}

	fmt.Printf("Saving file succeeded!")
	return nil
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid title or content")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
