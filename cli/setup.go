package cli

import (
	"fmt"

	"devely.com/mob/utils"
)

func Setup() {
	utils.ClearTerminal()
	fmt.Println("📦 Installing dependencies...")
	utils.PrintEmptyLine()
	fmt.Println("🔧 Installing FlutterGen")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"pub", "global", "activate", "flutter_gen"})
	utils.PrintEmptyLine()
	fmt.Println("🔧 Installing Rename")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"pub", "global", "activate", "rename"})
	utils.PrintEmptyLine()
	if !utils.IsWindows() {
		fmt.Println("🔧 Installing XcodeGen")
		utils.PrintEmptyLine()
		utils.RunCommandWithArgs("brew", []string{"install", "xcodegen"})
		utils.PrintEmptyLine()
	}
	fmt.Println("🔧 Installing Mason")
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("flutter", []string{"pub", "global", "activate", "mason_cli"})
	utils.PrintEmptyLine()
	utils.RunCommandWithArgs("mason", []string{"get"})
	utils.PrintEmptyLine()
	fmt.Println("🚀 Running flutter pub get...")
	utils.RunCommandWithArgs("flutter", []string{"pub", "get"})
}
