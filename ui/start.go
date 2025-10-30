package ui

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/niljimeno/gopher/tcp"
)

const (
	MODE_MENU = 0
	MODE_TXT  = 1
)

type ui_ struct {
	Mode   int
	Screen tcell.Screen
}

func (ui *ui_) Draw() {
	ui.Screen.Sync()
	ui.ShowContent()
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

func splitEvery(s string, n int) []string {
	var result []string
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	return result
}

func (ui ui_) ShowContent() {
	w, h := ui.Screen.Size()
	ui.Screen.Clear()
	// style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue.TrueColor()).Background(tcell.ColorWhite)

	var row int
	toPrint := tcp.Dial("gopher.quux.org:70", "")

	for _, v := range toPrint {
		if row >= h {
			break
		}

		if len(v.Content) <= w {
			emitStr(ui.Screen, 0, row, tcell.StyleDefault, v.Content)
			row++
			continue
		}

		chop := splitEvery(v.Content, w)
		for _, c := range chop {
			emitStr(ui.Screen, 0, row, tcell.StyleDefault, c)
			row++
		}

	}

	ui.Screen.Show()
}

func Start() error {
	ui := ui_{}
	var err error
	ui.Screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}

	if err := ui.Screen.Init(); err != nil {
		return err
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	ui.Screen.SetStyle(defStyle)

	ui.ShowContent()

	for {
		switch ev := ui.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			ui.Draw()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape ||
				ev.Rune() == 'q' {
				ui.Screen.Fini()
				os.Exit(0)
			}
		}
	}
}
