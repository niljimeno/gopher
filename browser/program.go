package browser

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/niljimeno/gopher/tcp"
	"github.com/niljimeno/gopher/types"
)

type program_ struct {
	Screen  tcell.Screen
	Buffers []buffer
	State   state
}

type buffer []tcp.Message

var emptyBuffer = buffer{{
	Type:    types.Information,
	Content: "Empty buffer",
}}

func (p *program_) mainLoop() {
	switch ev := p.Screen.PollEvent().(type) {
	case *tcell.EventResize:
		p.Draw()
	case *tcell.EventKey:
		p.HandleInput(ev)
	}
}

func (p *program_) HandleInput(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEscape:
		p.Quit()
	case tcell.KeyCtrlK:
		p.KillBuffer()
	}

	switch ev.Rune() {
	case 'q':
		p.Quit()
	}
}

func (p *program_) Buffer() *buffer {
	return &p.Buffers[p.State.BufferIndex]
}

func (p *program_) CreateBuffer(b buffer) {
	p.Buffers = append(p.Buffers, b)
	p.State.BufferIndex = len(p.Buffers) - 1
	p.Screen.PostEvent(&tcell.EventResize{})
}

func (p *program_) KillBuffer() {
	i := p.State.BufferIndex
	p.Buffers = append(p.Buffers[:i], p.Buffers[i+1:]...)
	p.State.BufferIndex--

	if p.State.BufferIndex <= 0 {
		p.CreateBuffer(emptyBuffer)
	}
	p.Screen.PostEvent(&tcell.EventResize{})
}

func (p *program_) Draw() {
	p.Screen.Sync()
	p.ShowScreen()
}

func (p *program_) LoadPage(url, route string) {
	p.State.Mode = LOADING
	go func() {
		buf := tcp.Dial(url, route)
		p.State.Mode = READING
		p.CreateBuffer(buf)
	}()
}

func (p program_) Quit() {
	p.Screen.Fini()
	os.Exit(0)
}
