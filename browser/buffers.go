package browser

import (
	"github.com/niljimeno/gopher/tcp"
)

type buffer struct {
	Content []tcp.Message
	Cursor  cursor
	Scroll  int
}

type cursor struct {
	Line int
}

func (b *buffer) moveCursor(step int) {
	prev := b.Cursor.Line
	b.Cursor.Line += step
	if 0 > b.Cursor.Line || b.Cursor.Line > len(b.Content)-1 {
		b.Cursor.Line = prev
	}
}
