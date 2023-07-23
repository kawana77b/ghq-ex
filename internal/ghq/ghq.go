package ghq

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kawana77b/ghq-ex/internal/util"
)

type Ghq struct {
	root  string
	repos []*Repository
}

func NewGhq(root string, repos []*Repository) *Ghq {
	return &Ghq{
		root:  filepath.Clean(root),
		repos: repos,
	}
}

func (g *Ghq) Root() string {
	return g.root
}

func (g *Ghq) Repos() []*Repository {
	return g.repos
}

func (g *Ghq) Names() []string {
	names := []string{}
	for _, repo := range g.repos {
		names = append(names, repo.Name())
	}

	return names
}

func (g *Ghq) Paths() []string {
	paths := []string{}
	for _, repo := range g.repos {
		paths = append(paths, repo.Path())
	}

	return paths
}

func (g *Ghq) Find(name string) *Repository {
	for _, repo := range g.repos {
		if repo.Name() == name {
			return repo
		}
	}

	return nil
}

func (g *Ghq) Cleanup() error {
	dirs := []string{}
	// root直下のディレクトリを全て取得する
	err := filepath.Walk(g.root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		if path == g.root {
			return nil
		}

		for _, repo := range g.repos {
			if repo.Contains(path) {
				return nil
			}
		}

		dirs = append(dirs, path)

		return nil
	})

	if err != nil {
		return err
	}

	// dirsをディレクトリ階層が深い順にソートする
	dirs = util.SortByPathDepth(dirs)
	for _, dir := range dirs {
		// ファイルおよびディレクトリが空であれば削除する
		di, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		if len(di) > 0 {
			continue
		}

		if err := os.Remove(dir); err != nil {
			return err
		}
	}

	return nil
}

func (g *Ghq) ChoiceRepoNameByPeco() (string, error) {
	if len(g.repos) == 0 {
		return "", nil
	}

	execPeco := func(buf []byte) string {
		cmd := exec.Command("peco")
		cmd.Stdin = bytes.NewBuffer(buf)

		out, err := cmd.Output()
		if err != nil {
			return ""
		}

		return strings.TrimSpace(string(out))
	}

	name := execPeco([]byte(strings.Join(g.Names(), "\n")))
	name = strings.ReplaceAll(strings.TrimSpace(name), "\n", "")
	return name, nil
}

func (g *Ghq) String() string {
	var value string = ""
	value += fmt.Sprintf("root: %s\n", g.root)

	for _, repo := range g.repos {
		v := fmt.Sprintf("%#v\n", repo)
		value += v
	}

	return value
}
