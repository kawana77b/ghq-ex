package util

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// パスの深さでソートする
func SortByPathDepth(paths []string) []string {
	depths := map[string]int{}
	for _, path := range paths {
		depths[path] = strings.Count(path, string(os.PathSeparator))
	}

	sort.Slice(paths, func(i, j int) bool {
		return depths[paths[i]] > depths[paths[j]]
	})

	return paths
}

// URLを規定ブラウザで開く.
// Linuxの場合、xdg-openコマンドを使用する. GUI環境である必要がある
func OpenURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		err = errors.New("unsupported platform")
	}

	return err
}

// 指定したディレクトリ以下のファイルを再帰的に取得する.
// ignoreHiddenDirがtrueの場合、隠しディレクトリは無視する.
// ignoreDotGitがtrueの場合、.gitディレクトリは無視する. これはignoreHiddenDirでも無視される、内容のため、falseの場合のみ効果がある
func GetFiles(dir string, ignoreHiddenDir bool, ignoreDotGit bool) ([]string, error) {
	files := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if ignoreHiddenDir && info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if ignoreDotGit && info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
