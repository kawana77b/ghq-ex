package cmd

import (
	"errors"
	"fmt"

	"github.com/Songmu/prompter"
	"github.com/kawana77b/ghq-ex/cmd/cmdutil"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the repository given as a selection or argument",
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
	g, err := cmdutil.MustGetGhq()
	if err != nil {
		return err
	}

	name, err := cmdutil.GetRepositoryName(g, args)
	if err != nil {
		return err
	}

	repo := g.Find(name)
	if repo == nil {
		return errors.New("repository not found")
	}

	fmt.Printf("You are trying to DELETE this local repository:\n")
	fmt.Println(aurora.Sprintf(aurora.Cyan("\t\t%s\n"), repo.Name()))
	isYes := prompter.YN("Are you sure?", false)
	if !isYes {
		return nil
	}

	if err := repo.Remove(); err != nil {
		return err
	}

	if err := g.Cleanup(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Remove %s\n", repo.Name())

	return nil
}
