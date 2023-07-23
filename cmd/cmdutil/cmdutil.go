package cmdutil

import (
	"errors"

	"github.com/kawana77b/ghq-ex/internal/ghq"
)

func MustGetGhq() (*ghq.Ghq, error) {
	ghqCmd := ghq.NewGhqCommand()
	if ghqCmd == nil {
		return nil, errors.New("ghq command not found")
	}

	g, err := ghqCmd.CreateGhq()
	if err != nil {
		return nil, err
	}

	if g == nil {
		return nil, errors.New("ghq command not found")
	}

	return g, nil
}

func GetRepositoryName(g *ghq.Ghq, args []string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	name, err := g.ChoiceRepoNameByPeco()
	if err != nil {
		return "", err
	}

	return name, nil
}
