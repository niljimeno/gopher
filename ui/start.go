package ui

import (
	"github.com/gdamore/tcell/v2"
)

func newProgram() (program_, error) {
	program := program_{Chan: make(chan string)}

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

func Start() error {
	program, err := newProgram()
	if err != nil {
		return err
	}

	program.LoadPage("gopher.quux.org:70", "")

	for {
		program.mainLoop()
	}
}
