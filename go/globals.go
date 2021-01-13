package main

import (
	"log"

	"github.com/gdamore/tcell"
)

const (
	boardXOffset = 4
	boardYOffset = 2
)

type (
	Board struct {
		boardsIndex  int
		width        int
		height       int
		colors       [][]tcell.Color
		rotation     [][]int
		dropDistance int
		fullLinesY   []bool
	}

	// View is the display engine
	View struct {
	}
)

var (
	logger *log.Logger
	screen tcell.Screen
	view   *View
)
