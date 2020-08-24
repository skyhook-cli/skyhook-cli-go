package model

import (
	"bufio"
	"fmt"
	"strings"

	. "github.com/logrusorgru/aurora"
)

type Prompt struct {
	Name, Question, Response, Default string
	Choices                           []string
}

func (p *Prompt) PrintPrompt() {
	s := p.Question
	if len(p.Choices) > 0 {
		s = fmt.Sprintf("%v\nchoices: [%v]", s, Blue(strings.Join(p.Choices, ", ")))
	}
	if p.Default != "" {
		s = fmt.Sprintf("%v\ndefault: %v", s, Green(p.Default))
	}

	fmt.Printf("%v %v\n%v ", Green("?"), s, Blue("->"))
}

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
	} else {
		return true
	}
}

func (p *Prompt) ReadResponse(r *bufio.Reader) {

	text := readHelper(r)

	if !p.ValidateResponse(text) {
		fmt.Println(Red("invalid response"))
	} else {
		p.Response = text
	}
}

func readHelper(r *bufio.Reader) string {
	text, ok := r.ReadString('\n')

	if ok == nil {
		text = strings.TrimSpace(text)
	} else {
		text = ""
	}

	return text
}
