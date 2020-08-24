package model

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/skyhook-cli/skyhook-cli-go/util"
)

/*
Prompt represents 1 question to be asked in the terminal, with any defaults, options, and the response
*/
type Prompt struct {
	Name, Question, Response, Default string
	Choices                           []string
}

/*
PrintPrompt prints the prompt to the terminal
*/
func (p *Prompt) PrintPrompt() {
	s := p.Question
	if len(p.Choices) > 0 {
		s = fmt.Sprintf("%v\nchoices: [%v]", s, aurora.Blue(strings.Join(p.Choices, ", ")))
	}
	if p.Default != "" {
		s = fmt.Sprintf("%v\ndefault: %v", s, aurora.Green(p.Default))
	}

	fmt.Printf("%v %v\n%v ", aurora.Green("?"), s, aurora.Blue("->"))
}

/*
ValidateResponse checks that the user's response is not empty and is a valid choice, if applicable
*/
func (p *Prompt) ValidateResponse(response string) bool {
	if len(response) == 0 {
		return false
	}
	if len(p.Choices) > 0 {
		for _, v := range p.Choices {
			if response == v {
				return true
			}
		}
		return false
	}
	return true
}

/*
ReadResponse reads the user input into the Response field of the prompt
*/
func (p *Prompt) ReadResponse(r *bufio.Reader) {

	text := inputOrDefault(util.ReadText(r), p.Default)
	isValid := p.ValidateResponse(text)

	for !isValid {
		fmt.Println(aurora.Red("invalid response"))
		p.PrintPrompt()

		text = inputOrDefault(util.ReadText(r), p.Default)
		isValid = p.ValidateResponse(text)
	}

	p.Response = text
}

func inputOrDefault(input string, defaultChoice string) string {
	if len(input) == 0 {
		return defaultChoice
	}
	return input
}
