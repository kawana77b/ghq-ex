package ghq

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Repository struct {
	name string
	path string
}

func NewRepository(name, path string) *Repository {
	return &Repository{
		name: name,
		path: filepath.Clean(path),
	}
}

func (r *Repository) Name() string {
	return r.name
}

func (r *Repository) Path() string {
	return r.path
}

func (r *Repository) Contains(path string) bool {
	p := filepath.Clean(path)
	return strings.HasPrefix(p, r.path)
}

func (r *Repository) Exists() bool {
	if _, err := os.Stat(r.path); err != nil {
		return false
	}

	return true
}

func (r *Repository) Remove() error {
	if !r.Exists() {
		return errors.New("repository not found")
	}

	// path傘下の全てのファイルパーミッションを777にする
	err := filepath.Walk(r.path, func(path string, info os.FileInfo, err error) error {
		if err := os.Chmod(path, 0777); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return os.RemoveAll(r.path)
}

func (r *Repository) RemoteUrl() string {
	// 現在のワーキングディレクトリを取得
	wd, err := os.Getwd()
	defer os.Chdir(wd)

	if err != nil {
		return ""
	}

	// このリポジトリに移動する
	if err := os.Chdir(r.path); err != nil {
		return ""
	}

	bytes, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return ""
	}

	url := strings.TrimSpace(string(bytes))
	if strings.HasPrefix(url, "git@") {
		url = strings.Replace(url, ":", "/", 1)
		url = strings.Replace(url, "git@", "https://", 1)
	}

	return url
}

// リポジトリの最新のタグ名を取得する. タグがない場合は空文字を返す
func (r *Repository) GetLatestTagName() string {
	if !r.Exists() {
		return ""
	}

	// 現在のワーキングディレクトリを取得
	wd, err := os.Getwd()
	defer os.Chdir(wd)
	if err != nil {
		return ""
	}

	// このリポジトリに移動する
	if err := os.Chdir(r.path); err != nil {
		return ""
	}

	bytes, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(bytes))
}
