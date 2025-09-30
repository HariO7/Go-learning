package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("\nThe todo you added was:%v\n ", todo.Text)
}

func (todo Todo) Save() error {

	fileName := "todo.json"

	data, err := json.Marshal(todo)

	if err != nil {
		fmt.Printf("\n Error occured while decoding:%v", err)
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

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid title or content")
	}
	return Todo{
		Text: text,
	}, nil
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
