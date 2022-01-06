package cli

import (
	"fmt"
	"os"

	"github.com/Devely-Team/mob/utils"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
)

func Module() {
	utils.ClearTerminal()
	choices := []string{
		"Feature",
		"Command",
	}
	sp := selection.New("O que deseja criar?", selection.Choices(choices))
	sp.FilterPrompt = ""
	sp.FilterPlaceholder = "Selecione ou digite o item desejado"

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	switch choice.Index {
	case 0:
		GenerateFeature()
	case 1:
		GenerateCommand()
	default:
		os.Exit(0)
	}
}

func GenerateFeature() {
	utils.ClearTerminal()
	utils.PrintEmptyLine()
	input := textinput.New("Digite um nome válido (lowercase, sem espaços): ")

	input.Validate = IsValidDartName
	input.CharLimit = 15

	name, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	utils.ClearTerminal()
	fmt.Println("📦 Generating feature...")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("mason", []string{"make", "feature", "-o", "lib/features/" + name, "--name", name, "--on-conflict", "prompt"})
}

func GenerateCommand() {
	utils.ClearTerminal()
	utils.PrintEmptyLine()
	input := textinput.New("Digite um nome válido (lowercase, sem espaços): ")

	input.Validate = IsValidDartName
	input.CharLimit = 15

	name, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	utils.ClearTerminal()
	fmt.Println("📦 Generating command...")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("mason", []string{"make", "command", "-o", "lib/domain/commands/" + name, "--name", name, "--on-conflict", "prompt"})
}

func IsValidDartName(name string) bool {
	if len(name) < 1 {
		return false
	}

	if name[0] >= '0' && name[0] <= '9' {
		return false
	}

	for i := 0; i < len(name); i++ {
		if name[i] >= 'a' && name[i] <= 'z' {
			continue
		}

		if name[i] >= '0' && name[i] <= '9' {
			continue
		}

		if name[i] == '_' {
			continue
		}

		return false
	}

	return true
}
