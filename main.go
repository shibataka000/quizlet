package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shibataka000/quizlet/internal/quizlet"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		word, skip, err := quizlet.Parse(text)
		if err != nil {
			fmt.Printf("[WARN] %v\n", err)
			continue
		}
		if skip {
			continue
		}
		switch word.Kind {
		case "expression":
			fmt.Printf("%s\t%s\n", word.English, word.Japanese)
		default:
			fmt.Printf("%s（%s）\t%s\n", word.English, word.Japanese, word.Description)
		}
	}
}
