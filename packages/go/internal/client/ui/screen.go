package ui

import (
	"fmt"
	. "github.com/brickshot/roadtrip-v2/internal/client"
)

var esc = "\033"

type Feature struct {
	c string
	z int
}

type Screen struct {
	Width  int
	Height int
}

func cls() {
	fmt.Print("\033[2J")
}

func home() {
	fmt.Printf("%s[H", esc)
}

func moveToColumn(c int) {
	fmt.Printf("%s[%dG", esc, c)
}

func moveTo(x int, y int) {
	fmt.Printf("%s[%d;%dH", esc, y, x)
}

func color(fg int, bg int) {
	fmt.Printf("%s[%d;%dm", esc, fg, bg)
}

type RenderData struct {
	CurrentCharacter CurrentCharacter
}

func Render(data RenderData) {
	cls()
	moveTo(0, 0)
	m := data.CurrentCharacter.Car.Route.Map
Outer:
	for r := 0; r < m.H; r++ {
		for c := 0; c < m.W; c++ {
			i := r*m.W + c
			if i >= len(m.Image) {
				break Outer
			}
			fmt.Printf("%c", m.Image[i])
		}
		fmt.Printf("\n")
	}

	r := data.CurrentCharacter.Car.Route
	i := -1
	var cn Node
	var cr Road
	// this logic needs a rewrite bc it's being used to both render and get current node and current road
	// we don't want to have to rerender the road every time just to get this stuff
	for _, w := range r.Ways {
		for _, n := range w.Nodes {
			i++
			for _, f := range n.Features {
				if f.Road.Glyph != "" {
					moveTo(n.X, n.Y)
					if data.CurrentCharacter.Car.Location.Index == i {
						fmt.Printf("☺")
						cn = n
						cr = f.Road
					} else {
						fmt.Printf("%v", f.Road.Glyph)
					}
				}
			}
		}
	}
	moveTo(1, 0)
	fmt.Printf("Route: %s", r.Name)
	moveTo(1, 24)
	fmt.Printf("In %s on %s going %v MPH", cn.State.Name, cr.Name, data.CurrentCharacter.Car.MPH)
}
