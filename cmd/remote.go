/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kawana77b/ghq-ex/cmd/cmdutil"
	"github.com/kawana77b/ghq-ex/internal/util"
	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Open the remote URL in the specified browser",
	Long: `Open the remote URL in the specified browser.

This command basically requires the remote URL to be registered as http, https.
	`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1)),
	RunE: runRemote,
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runRemote(cmd *cobra.Command, args []string) error {
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

	url := repo.RemoteUrl()
	if url == "" {
		return errors.New("remote url not found")
	}

	for _, prefix := range []string{"https://", "http://"} {
		if strings.Contains(url, prefix) {
			fmt.Printf("Open: %s\n", url)

			if err := util.OpenURL(url); err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("unsupported remote url")
}
