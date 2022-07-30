package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	ball := Ball{
		X:      1,
		Y:      1,
		Xspeed: 1,
		Yspeed: 1,
	}

	game := Game{
		Screen: s,
		Ball:   ball,
	}

	go game.Run()

	for {
		switch event := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
