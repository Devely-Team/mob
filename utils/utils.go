package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(command string) error {
	return RunCommandWithArgs(command, []string{})
}

func RunCommandWithArgs(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func IsWindows() bool {
	return os.PathSeparator == '\\'
}

func ClearTerminal() {
	if IsWindows() {
		RunCommand("cls")
	} else {
		RunCommand("clear")
	}
}

func PrintEmptyLine() {
	fmt.Println("")
}
