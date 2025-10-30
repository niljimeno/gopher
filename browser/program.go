package browser

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/niljimeno/gopher/tcp"
)

type program_ struct {
	Screen  tcell.Screen
	Buffers []buffer
	State   state
}

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
	case 'k':
		p.Buffer().moveCursor(-1)
		p.RefreshScreen()
	case 'j':
		p.Buffer().moveCursor(+1)
		p.RefreshScreen()
	}
}

func (p *program_) RefreshScreen() {
	p.Screen.PostEvent(&tcell.EventResize{})
}

func (p *program_) Buffer() *buffer {
	return &p.Buffers[p.State.BufferIndex]
}

func (p *program_) CreateBuffer(m []tcp.Message) {
	p.Buffers = append(p.Buffers, buffer{
		Content: m,
	})
	p.State.BufferIndex = len(p.Buffers) - 1
	p.Screen.PostEvent(&tcell.EventResize{})
}

func (p *program_) CreateEmptyBuffer() {
	p.CreateBuffer([]tcp.Message{{Content: "Empty buffer"}})
}

func (p *program_) KillBuffer() {
	i := p.State.BufferIndex
	p.Buffers = append(p.Buffers[:i], p.Buffers[i+1:]...)
	p.State.BufferIndex--

	if p.State.BufferIndex <= 0 {
		p.CreateEmptyBuffer()
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
