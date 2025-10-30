package browser

import (
	"github.com/gdamore/tcell/v2"
)

func Start() error {
	program, err := newProgram()
	if err != nil {
		return err
	}

	program.LoadPage("gopher.meulie.net:70", "")

	for {
		program.mainLoop()
	}
}

func newProgram() (program_, error) {
	program := program_{}

	var err error
	program.Screen, err = tcell.NewScreen()
	if err != nil {
		return program_{}, err
	}

	if err := program.Screen.Init(); err != nil {
		return program_{}, err
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	program.Screen.SetStyle(defStyle)

	return program, nil
}
