package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"

	"github.com/evertras/twitchdash/pkg/stream"
)

func StartLayout(g *gocui.Gui, info stream.Information, viewers <-chan int) func(g *gocui.Gui) error {
	go func() {
		for viewerCount := range viewers {
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("chat")

				if err != nil {
					return err
				}

				v.Subtitle = fmt.Sprintf("%d viewers", viewerCount)
				return nil
			})
		}
	}()

	return func(g *gocui.Gui) error {
		return viewChat(info)(g)
	}
}
