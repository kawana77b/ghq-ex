package cmd

import (
	"errors"
	"fmt"

	"github.com/kawana77b/ghq-ex/internal/ghq"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Outputs the full path to the repository given as a selection or argument",
	Long:  ``,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	RunE:  runShow,
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runShow(cmd *cobra.Command, args []string) error {
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

	fmt.Println(repo.Path())

	return nil
}
