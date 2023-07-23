package util

import (
	"os"
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
