package ghq

import (
	"errors"
	"os"
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
