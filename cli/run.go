package cli

import (
	"fmt"
	"os"

	"devely.com/mob/utils"
	"github.com/erikgeiser/promptkit/selection"
)

func Run() {
	utils.ClearTerminal()
	utils.PrintEmptyLine()
	choices := []string{
		"Projeto",
		"Build Runner Watch",
	}

	sp := selection.New("O que gostaria de rodar?", selection.Choices(choices))
	sp.FilterPrompt = ""
	sp.FilterPlaceholder = "Selecione ou digite o item desejado"
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	utils.ClearTerminal()
	switch choice.Index {
	case 0:
		utils.PrintEmptyLine()
		fmt.Println("📦 Running project...")
		utils.PrintEmptyLine()
		utils.RunCommandWithArgs("flutter", []string{"run"})
	case 1:
		utils.PrintEmptyLine()
		fmt.Println("📦 Running build runner watch...")
		utils.PrintEmptyLine()
		utils.RunCommandWithArgs("flutter", []string{"pub", "run", "build_runner", "watch"})
	default:
		os.Exit(0)
	}
}
