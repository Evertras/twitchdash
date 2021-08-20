package ui

import (
	"github.com/awesome-gocui/gocui"

	"github.com/evertras/twitchdash/pkg/stream"
)

func Layout(info stream.Information) func (g *gocui.Gui) error {
	return viewChat(info)
}
