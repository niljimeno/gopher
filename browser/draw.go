package browser

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func (p program_) ShowScreen() {
	w, h := p.Screen.Size()
	p.Screen.Clear()

	switch p.State.Mode {
	case LOADING:
		loadingText := "Loading..."
		emitStr(
			p.Screen,
			(w-len(loadingText))/2, h/2-1,
			tcell.StyleDefault,
			loadingText,
		)

	case READING:
		var row int

		for _, v := range *p.Buffer() {
			if row >= h {
				break
			}

			if len(v.Content) <= w {
				emitStr(p.Screen, 0, row, tcell.StyleDefault, v.Content)
				row++
				continue
			}

			chop := splitEvery(v.Content, w)
			for _, c := range chop {
				emitStr(p.Screen, 0, row, tcell.StyleDefault, c)
				row++
			}
		}
	}

	p.Screen.Show()
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
		end := min(i+n, len(s))
		result = append(result, s[i:end])
	}
	return result
}
