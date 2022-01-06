package cli

import (
	"fmt"
	"os"

	"github.com/Devely-Team/mob/utils"
	"github.com/erikgeiser/promptkit/selection"
)

func Build() {
	utils.ClearTerminal()
	choices := []string{
		"Android Release (APK)",
		"iOS Release (IPA)",
		"All",
	}

	sp := selection.New("O que gostaria de buildar?", selection.Choices(choices))
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
		BuildAndroidRelease()
	case 1:
		BuildIosRelease()
	case 2:
		BuildAll()
	default:
		os.Exit(0)
	}
}

func BuildAndroidRelease() {
	utils.PrintEmptyLine()
	fmt.Println("📦 Building Android Release (APK)")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"build", "apk", "--release"})
}

func BuildIosRelease() {
	utils.PrintEmptyLine()
	fmt.Println("📦 Building iOS Release (IPA)")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"build", "ios"})
}

func BuildAll() {
	utils.PrintEmptyLine()
	fmt.Println("📦 Building All")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"build"})
}
