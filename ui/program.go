package ui

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/niljimeno/gopher/tcp"
)

type program_ struct {
	Buffer []tcp.Message
	Mode   int
	Screen tcell.Screen
	Chan   chan string
}

func (p *program_) mainLoop() {
	switch ev := p.Screen.PollEvent().(type) {
	case *tcell.EventResize:
		p.Draw()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape ||
			ev.Rune() == 'q' {
			p.Screen.Fini()
			os.Exit(0)
		}
	}
}

func (p *program_) Draw() {
	p.Screen.Sync()
	p.ShowScreen()
}

func (p *program_) LoadPage(url, route string) {
	p.Mode = LOADING
	go func() {
		time.Sleep(time.Second)

		p.Buffer = tcp.Dial(url, route)
		p.Mode = READING
		p.Screen.PostEvent(&tcell.EventResize{})
	}()
}
