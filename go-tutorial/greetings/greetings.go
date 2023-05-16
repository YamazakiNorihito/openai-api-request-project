package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// :=演算子は1行で変数を宣言および初期化するためのショートカットです
	// var message string = "Hello, " + name + "!"

	// 変数の宣言と初期化を2行に分けたもの
	// message := "Hello, " + name + "!"
    return message, nil
}