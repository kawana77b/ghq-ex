package cmd

import (
	"errors"
	"fmt"

	"github.com/kawana77b/ghq-ex/internal/ghq"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up empty folders in the ghq folder",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(0)),
	RunE:  runCleanup,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCleanup(cmd *cobra.Command, args []string) error {
	ghqCmd := ghq.NewGhqCommand()
	if cmd == nil {
		return errors.New("ghq command not found")
	}

	g, err := ghqCmd.CreateGhq()
	if err != nil {
		return err
	}

	if err := g.Cleanup(); err != nil {
		return err
	}

	fmt.Println("Cleanup done.")

	return nil
}
