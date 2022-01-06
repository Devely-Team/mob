package cli

import (
	"fmt"

	"github.com/Devely-Team/mob/utils"
)

func Test() {
	utils.ClearTerminal()
	utils.PrintEmptyLine()
	fmt.Println("📦 Running tests...")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"test", ".", "--no-pub", "--coverage", "--suppress-analytics"})
}
