package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

// NewView creates a new view
func NewView() {
	var err error

	screen, err = tcell.NewScreen()
	if err != nil {
		logger.Fatal("NewScreen error:", err)
	}
	err = screen.Init()
	if err != nil {
		logger.Fatal("screen Init error:", err)
	}

	screen.Clear()
	rand.Seed(time.Now().UnixNano())

	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			styleBoarder := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
			if rand.Float32() > 0.2 {
				styleBoarder = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)
			}
			screen.SetContent(x, y, ' ', nil, styleBoarder)
		}
	}
	screen.Show()

	view = &View{}
}

// Stop stops the view
func (view *View) Stop() {
	logger.Println("View Stop start")

	screen.Fini()

	logger.Println("View Stop end")
}

func (view *View) drawTexts() {
	/*
		xOffset := boardXOffset + board.width*2 + 8
		yOffset := boardYOffset + 7

		view.drawText(xOffset, yOffset, "SCORE:", tcell.ColorLightGray, tcell.ColorDarkBlue)
	*/
}

func (view *View) drawText(x int, y int, text string, fg tcell.Color, bg tcell.Color) {
	style := tcell.StyleDefault.Foreground(fg).Background(bg)
	for index, char := range text {
		screen.SetContent(x+index, y, rune(char), nil, style)
	}
}
