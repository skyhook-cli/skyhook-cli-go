package util

import (
	"bufio"
	"strings"
)

/*
ReadText checks for errors when reading user input from stdin
*/
func ReadText(r *bufio.Reader) string {
	text, ok := r.ReadString('\n')

	if ok == nil {
		text = strings.TrimSpace(text)
	} else {
		text = ""
	}

	return text
}
