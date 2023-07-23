package util

import (
	"errors"
	"os"
	"os/exec"
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
