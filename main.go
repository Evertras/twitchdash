package main

import (
	"errors"
	"log"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/evertras/twitchdash/pkg/stream"
	"github.com/evertras/twitchdash/pkg/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	info := stream.Information{
		Title: "My awesome stream",
	}

	viewers := make(chan int)

	go func() {
		counter := 0
		for {
			counter++
			viewers <- counter
			time.Sleep(time.Millisecond * 100)
		}
	}()

	g.SetManagerFunc(ui.StartLayout(g, info, viewers))

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
