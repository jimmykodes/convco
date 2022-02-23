package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

var (
	types = []string{
		"feat",
		"fix",
		"docs",
		"style",
		"refactor",
		"perf",
		"test",
		"build",
		"ci",
		"chore",
		"revert",
	}
	helpText = `
feat       A new feature
fix        A bug fix
docs       Documentation only changes
style      Style changes (white space, formatting, missing semi-colons, etc)
refactor   A code change that does not fix a bug or implement a feature
perf       A code change to improve performance
test       Adding missing tests or correcting exising ones
build      Changes that affect build system (example scopes: pypi, go.mod)
ci         Changes to CI files (example scopes: Travis, Circle, Actions)
chore      Other changes that don't modify src or test files
revert     Reverts a previous commit
`
	qs = []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Choose a type:",
				Options: types,
				Default: "feat",
				Help:    helpText,
			},
			Validate: survey.Required,
		},
		{
			Name:   "scope",
			Prompt: &survey.Input{Message: "Scope (optional):"},
		},
		{
			Name:     "description",
			Prompt:   &survey.Input{Message: "Description:"},
			Validate: survey.Required,
		},
	}
)

func main() {
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		fmt.Println(helpText)
		return
	}

	answers := struct {
		Type        string
		Scope       string
		Description string
	}{}
	err := survey.Ask(qs, &answers)
	if errors.Is(err, terminal.InterruptErr) {
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s", err)
		return
	}
	var sb strings.Builder
	sb.WriteString(answers.Type)
	if answers.Scope != "" {
		fmt.Fprintf(&sb, "(%s)", answers.Scope)
	}
	sb.WriteString(": ")
	sb.WriteString(answers.Description)
	fmt.Println(sb.String())
}
