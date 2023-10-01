package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kawana77b/ghq-ex/cmd/cmdutil"
	"github.com/kawana77b/ghq-ex/internal/util"
	"github.com/spf13/cobra"
)

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "Create a zip file to upload to Go's package registry",
	Long: `Create a zip file to upload to Go's package registry.

This command can be used for any repository, but is targeted at Golang projects.
More information can be found at the official site below:

https://go.dev/ref/mod#zip-files
`,
	RunE: runZip,
}

func init() {
	rootCmd.AddCommand(zipCmd)
}

func runZip(cmd *cobra.Command, args []string) error {
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

	// .gitディレクトリを除いた全てのファイルを取得
	files, err := util.GetFiles(repo.Path(), false, true)
	if err != nil {
		return err
	}

	getVersion := func() string {
		version := repo.GetLatestTagName()
		if len(version) == 0 {
			return "v0.0.1"
		}

		return version
	}

	repoName := filepath.Base(repo.Path())
	version := getVersion()
	validRepoNameInGolang := fmt.Sprintf("%s@%s", repoName, version)

	// ビルダにファイルを追加
	builder := util.NewZipFileBuilder()
	for _, file := range files {
		formatToName := func(path string, root string) string {
			name := strings.Replace(path, root, "", 1)
			name = strings.ReplaceAll(name, "\\", "/") // Windows対応
			name = strings.TrimPrefix(name, "/")       // 先頭の / を削除

			// リポジトリ名をGOで有効なフォルダ名に置換する
			name = strings.Replace(name, repoName, validRepoNameInGolang, 1)

			return name
		}

		name := formatToName(file, g.Root())
		builder.AddFile(name, file)
	}

	// 現在のワークディレクトリを取得
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// zipファイルを作成
	fileName := fmt.Sprintf("%s.zip", validRepoNameInGolang)
	filePath := filepath.Join(wd, fileName)
	if err := builder.Create(filePath); err != nil {
		return err
	}

	fmt.Printf("created: %s\n", filePath)

	return nil
}
