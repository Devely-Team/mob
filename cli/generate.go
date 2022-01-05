package cli

import (
	"fmt"

	"devely.com/mob/utils"
)

func Generate() {
	utils.ClearTerminal()
	utils.PrintEmptyLine()
	fmt.Println("📦 Generating...")
	utils.PrintEmptyLine()
	fmt.Println("🚀 Running flutter pub get...")
	utils.RunCommandWithArgs("flutter", []string{"pub", "get"})
	utils.PrintEmptyLine()
	fmt.Println("🚀 Generating assets...")
	utils.RunCommandWithArgs("fluttergen", []string{"."})
	utils.PrintEmptyLine()
	fmt.Println("🚀 Generating i18n...")
	utils.RunCommandWithArgs("flutter", []string{"pub", "run", "fast_i18n"})
	utils.PrintEmptyLine()
	fmt.Println("🚀 Generating xcode project...")
	utils.RunCommandWithArgs("xcodegen", []string{"generate", "--spec", "ios/project.yml", "--project", "ios/"})
	utils.RunCommandWithArgs("pod", []string{"install", "--project-directory=ios"})
}
