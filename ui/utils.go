package ui

import (
	"os"

	"github.com/aymanbagabas/go-osc52/v2"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dlvhdr/gh-dash/v4/data"
	"github.com/dlvhdr/gh-dash/v4/ui/constants"
)

func oscCopy(text string) error {
	s := osc52.New(text)

	if _, ok := os.LookupEnv("TMUX"); ok {
		s = s.Tmux()
	} else if _, ok := os.LookupEnv("STY"); ok {
		s = s.Screen()
	}

	s.WriteTo(os.Stdout)
	return nil
}

type userFetchedMsg struct {
	user string
}

func fetchUser() tea.Msg {
	user, err := data.CurrentLoginName()
	if err != nil {
		return constants.ErrMsg{
			Err: err,
		}
	}

	return userFetchedMsg{
		user: user,
	}
}
