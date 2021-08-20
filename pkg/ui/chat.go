package ui

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"

	"github.com/evertras/twitchdash/pkg/stream"
)

func viewChat(info stream.Information) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		maxX, maxY := g.Size()
		if v, err := g.SetView("chat", 0, 0, maxX-1, maxY-1, 0); err != nil {
			v.Frame = true
			v.Title = info.Title
			v.Wrap = true
			v.Autoscroll = true
			v.FrameRunes = []rune{' ', ' '}
			if !errors.Is(err, gocui.ErrUnknownView) {
				return err
			}

			if _, err := g.SetCurrentView("chat"); err != nil {
				return err
			}

			for i := 0; i < 50; i++ {
				fmt.Fprintln(v, i)
				fmt.Fprintln(v, "\x1b[0;31mNevertras: \x1b[0;00mOk this is a really long line with some wrapping eventually maybe idk this is just a lot of typing I guess and here we go we needed more more more more more more more mor emore more more mroe more moermeorm eormeroemroerm emomroemroer")
			}
		}

		return nil
	}
}
