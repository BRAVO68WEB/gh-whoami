package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh whoami",
		Short: "gh whoami is a command line tool to check who you are logged in as on GitHub",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSearch()
		},
	}

	return cmd
}

func runSearch() error {
	args := []string{"api", "user", "--jq", ".login"}

	username, _, err := gh.Exec(args...)

	if err != nil {
		return err
	}

	usernameString := username.String()

	fmt.Printf("%s\n", usernameString)

	return nil
}

func main() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
