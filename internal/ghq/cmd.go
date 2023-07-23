package ghq

import (
	"os/exec"
	"path/filepath"
	"strings"
)

type GhqCommand struct {
	path string
}

func NewGhqCommand() *GhqCommand {
	path, err := exec.LookPath("ghq")
	if err != nil {
		return nil
	}

	return &GhqCommand{
		path: path,
	}
}

func (g *GhqCommand) Command(args ...string) *exec.Cmd {
	return exec.Command(g.path, args...)
}

func (g *GhqCommand) Root() (string, error) {
	bytes, err := g.Command("root").Output()
	if err != nil {
		return "", err
	}

	value := strings.TrimSpace(string(bytes))
	path := filepath.Clean(value)
	return path, nil
}

func (g *GhqCommand) List() ([]string, error) {
	bytes, err := g.Command("list").Output()
	if err != nil {
		return []string{}, err
	}

	value := strings.TrimSpace(string(bytes))
	if value == "" {
		return []string{}, nil
	}

	paths := strings.Split(value, "\n")
	return paths, nil
}

func (g *GhqCommand) CreateGhq() (*Ghq, error) {
	root, err := g.Root()
	if err != nil {
		return nil, err
	}

	names, err := g.List()
	if err != nil {
		return nil, err
	}

	repos := []*Repository{}
	for _, name := range names {
		repos = append(repos, NewRepository(name, filepath.Join(root, name)))
	}

	return NewGhq(root, repos), nil
}
