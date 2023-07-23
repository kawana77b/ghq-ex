package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the ghq repository all at once",
	Long: `Update the ghq repository all at once

This command is the same as:
    ghq list | ghq get --update --parallel`,
	Args: cobra.ExactArgs(0),
	RunE: runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runUpdate(cmd *cobra.Command, args []string) error {
	if _, err := exec.LookPath("ghq"); err != nil {
		return fmt.Errorf("ghq command not found: %s", err)
	}

	listCmd := exec.Command("ghq", "list")
	listOutput, err := listCmd.Output()
	if err != nil {
		return fmt.Errorf("error executing 'ghq list': %s", err)
	}

	repoPaths := strings.Split(string(listOutput), "\n")

	for _, path := range repoPaths {
		if path != "" {
			getCmd := exec.Command("ghq", "get", "--update", "--parallel", path)
			getCmd.Stdout = os.Stdout
			getCmd.Stderr = os.Stderr

			err := getCmd.Run()
			if err != nil {
				return fmt.Errorf("error executing 'ghq get --update --parallel %s': %s", path, err)
			}
		}
	}

	return nil
}
