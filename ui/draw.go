package ui

import "github.com/gdamore/tcell/v2"

func (p program_) ShowScreen() {
	w, h := p.Screen.Size()
	p.Screen.Clear()

	switch p.Mode {
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

		for _, v := range p.Buffer {
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
