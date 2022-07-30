package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Ball   Ball
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	g.Screen.SetStyle(defStyle)

	for {
		width, height := g.Screen.Size()
		g.Ball.CheckEdges(width, height)
		g.Screen.Clear()

		g.Ball.Update()

		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
