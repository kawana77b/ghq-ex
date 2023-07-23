package cmd

import (
	"errors"
	"fmt"

	"github.com/kawana77b/ghq-ex/internal/ghq"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	RunE:  runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runRemove(cmd *cobra.Command, args []string) error {
	ghqCmd := ghq.NewGhqCommand()
	if cmd == nil {
		return errors.New("ghq command not found")
	}

	g, err := ghqCmd.CreateGhq()
	if err != nil {
		return err
	}

	name := ""
	if len(args) > 0 {
		name = args[0]
	}

	if name == "" {
		name, err = g.ChoiceRepoNameByPeco()
		if err != nil {
			return err
		}
	}

	repo := g.Find(name)
	if repo == nil {
		return errors.New("repository not found")
	}

	if err := repo.Remove(); err != nil {
		return err
	}

	if err := g.Cleanup(); err != nil {
		return err
	}

	fmt.Printf("Remove %s\n", repo.Name())

	return nil
}
