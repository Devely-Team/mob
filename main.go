package main

import (
	"fmt"
	"os"

	"github.com/Devely-Team/mob/cli"
	"github.com/erikgeiser/promptkit/selection"
)

func main() {
	sp := selection.New("O que deseja fazer?", selection.Choices(getChoices()))
	sp.FilterPrompt = ""
	sp.FilterPlaceholder = "Selecione ou digite o item desejado"

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	switch choice.Index {
	case 0:
		cli.Setup()
	case 1:
		cli.Build()
	case 2:
		cli.Test()
	case 3:
		cli.Run()
	case 4:
		cli.Generate()
	case 5:
		cli.Module()
	default:
		os.Exit(0)
	}
}

func getChoices() []string {
	return []string{
		"Setup",
		"Build",
		"Test",
		"Run",
		"Generate",
		"Module",
	}
}
