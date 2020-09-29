package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

type FileData struct {
	User string `json:"user"`
	Passwd string `json:"passwd"`
}

func NoEchoPrompt(p string) ([]byte, error) {
	fmt.Printf(p)
	defer fmt.Printf("\n")
	return terminal.ReadPassword(0)
}

func main() {

	var readCmd = &cobra.Command{
		Use: "read [from file.ysnp]",
		Short: "Read password file",
		Long: "Read password file",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			keyToUse, err := NoEchoPrompt("key>")
			if err != nil {
				os.Exit(-1)
			}

			ReadFile(keyToUse, args[0])
		},
	}

	var writeCmd = &cobra.Command{
		Use: "write [to file.ysnp]",
		Short: "Write password file",
		Long: "Write password file",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			keyToUse, err := NoEchoPrompt("key>")
			if err != nil {
				os.Exit(-1)
			}

			usernameToStore, err := NoEchoPrompt("username >")
			if err != nil {
				os.Exit(-1)
			}

			passToStore, err := NoEchoPrompt("password >")
			if err != nil {
				os.Exit(-1)
			}

			WriteFile(keyToUse, FileData{string(usernameToStore), string(passToStore)}, args[0])
		},
	}

	var rootCmd = &cobra.Command{Use: "ysnp"}
	rootCmd.AddCommand(readCmd, writeCmd)
	rootCmd.Execute()
}

