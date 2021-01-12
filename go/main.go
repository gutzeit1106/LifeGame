package main

import (
	"time"
	"math/rand"
	"log"
	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = screen.Init(); err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	
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

	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev.(type) {
			case *tcell.EventKey:
				close(quit)
			}
		}
	}()
	<-quit
}