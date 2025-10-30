package browser

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/niljimeno/gopher/tcp"
	"github.com/niljimeno/gopher/types"
)

var alternateStyle = tcell.StyleDefault.
	Background(tcell.ColorWhite).
	Foreground(tcell.ColorBlack)

func (p program_) ShowScreen() {
	w, h := p.Screen.Size()
	p.Screen.Clear()
	switch p.State.Mode {
	case LOADING:
		p.ShowLoadingScreen(w, h)
	case READING:
		p.ShowPage(w, h)
	}

	p.Screen.Show()
}

func (p program_) ShowLoadingScreen(w, h int) {
	printCentered(p.Screen, "Loading...", w, h)
}

func printCentered(s tcell.Screen, text string, w, h int) {
	emitStr(
		s,
		(w-len(text))/2, h/2-1,
		tcell.StyleDefault,
		text)
}

func (p *program_) ShowPage(w, h int) {
	b := p.Buffer()

	var ptr = new(int)
	*ptr = -1

	if b.Cursor.Line < b.Scroll+5 {
		b.Scroll = b.Cursor.Line - 5
		if b.Scroll < 0 {
			b.Scroll = 0
		}
	} else if b.Cursor.Line >= b.Scroll+h-5 {
		b.Scroll = b.Cursor.Line - h + 5
	}

	start := b.Scroll

	for i, v := range b.Content[start:] {
		text := formattedMessage(v)

		var style tcell.Style
		if start+i == b.Cursor.Line {
			style = alternateStyle
		} else {
			style = tcell.StyleDefault
		}

		printMessage(p.Screen, style, ptr, text, w)
	}
}

func printMessage(s tcell.Screen, style tcell.Style, ptr *int, text string, maxWidth int) {
	*ptr++
	if len(text) <= maxWidth {
		spacing := strings.Repeat(" ", maxWidth-len(text))
		emitStr(s, 0, *ptr, style, text+spacing)
	} else {
		emitStr(s, 0, *ptr, style, text)
	}

	//if len(text) <= maxWidth {
	//	spacing := strings.Repeat(" ", maxWidth-len(text))
	//	emitStr(s, 0, *ptr, style, text+spacing)
	//} else {
	//	emitStr(s, 0, *ptr, style, text[:maxWidth])
	//	printMessage(s, style, ptr, text[maxWidth:], maxWidth)
	//}
}

func formattedMessage(m tcp.Message) string {
	switch m.Type {
	default:
		return "[?]" + m.Content
	case types.Information:
		return m.Content
	case types.SubMenu:
		return "[Submenu] " + m.Content
	}
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}
