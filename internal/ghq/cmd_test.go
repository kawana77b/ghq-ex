package ghq_test

import (
	"testing"

	"github.com/kawana77b/ghq-ex/internal/ghq"
)

// *GhqCommand.Rootのテスト
func Test_Root(t *testing.T) {
	ghq := ghq.NewGhqCommand()
	if ghq == nil {
		t.Fatal("ghq command not found")
	}

	path, err := ghq.Root()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(path)
}

func Test_List(t *testing.T) {
	ghq := ghq.NewGhqCommand()
	if ghq == nil {
		t.Fatal("ghq command not found")
	}

	paths, err := ghq.List()
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		t.Log(path)
	}
}

func Test_CreateGhq(t *testing.T) {
	cmd := ghq.NewGhqCommand()
	if cmd == nil {
		t.Fatal("ghq command not found")
	}

	ghq, err := cmd.CreateGhq()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ghq.String())
}
