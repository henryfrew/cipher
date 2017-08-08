package main

import (
	"fmt"

	"github.com/henryfrew/utils"
	c "github.com/skilstak/go-colors"
)

func main() {
	input := utils.Input(c.Y + "Would you like to encode or decode?\n> ")
	if input == "encode" {
		encode()
	} else if input == "decode" {
		decode()
	}
}

func encode() {
	message := utils.Input(c.Y + "Enter a message to be encoded.\n> ")
	morse := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}
	for i := 0; i < len(message); i++ {
		char := string(message[i])
		fmt.Print(c.B + morse[char] + " ")
	}
}

func decode() {
	message := utils.Input(c.Y + "Enter a message to be decoded.\n> ")
	/*morse := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}
	*/
	fmt.Print(message)
}
