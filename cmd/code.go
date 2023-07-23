package cmd

import (
	"errors"
	"os/exec"

	"github.com/kawana77b/ghq-ex/cmd/cmdutil"
	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Open selected or specified repository in VSCode",
	Long:  ``,
	RunE:  runCode,
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCode(cmd *cobra.Command, args []string) error {
	code, err := exec.LookPath("code")
	if err != nil {
		return errors.New("code command not found")
	}

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

	if err := exec.Command(code, repo.Path()).Run(); err != nil {
		return err
	}

	return nil
}
