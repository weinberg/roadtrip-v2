package ui

import (
	"fmt"
	. "github.com/brickshot/roadtrip-v2/internal/client"
	"math"
)

type Feature struct {
	c string
	z int
}

type Screen struct {
	Width  int
	Height int
}

type RenderData struct {
	CurrentCharacter CurrentCharacter
}

var esc = "\033"

func enterAlternateScreen() {
	fmt.Printf("%s[?1049h", esc)
}

func exitAlternateScreen() {
	fmt.Printf("%s[?1049l", esc)
	fmt.Println("Thanks for playing")
}

func cls() {
	fmt.Print("\x0c")
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

func (s Screen) Start() {
	enterAlternateScreen()
}

func (s Screen) Finish() {
	exitAlternateScreen()
}

func (s Screen) Render(data RenderData) {
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
	var ct Town
	// this logic needs a rewrite bc it's being used to both render and get current node and current road
	// we don't want to have to rerender the road every time just to get this stuff
	for _, w := range r.Ways {
		for _, n := range w.Nodes {
			i++
			drawGlyphs := []string{"", "", ""} // ordered layers - 0:road, 1:town, 2:car
			for _, f := range n.Features {
				if f.Road.Typename == "Road" {
					drawGlyphs[0] = string(f.Road.Glyph)
					if data.CurrentCharacter.Car.Location.Index == i {
						cr = f.Road
					}
				} else if f.Town.Typename == "Town" {
					drawGlyphs[1] = string(f.Road.Glyph)
					if data.CurrentCharacter.Car.Location.Index == i {
						ct = f.Town
					}
				}
			}
			if data.CurrentCharacter.Car.Location.Index == i {
				drawGlyphs[2] = "☺"
				cn = n
			}

			// reverse over glyphs to pick highest precedence glyph to draw
			for i := 2; i > -1; i-- {
				if drawGlyphs[i] != "" {
					moveTo(n.X, n.Y)
					fmt.Printf("%v", drawGlyphs[i])
					break
				}
			}
		}
	}

	titleX := int(math.Max(1, math.Floor(float64(s.Width)/2-float64(len(r.Name))/2)))
	moveTo(titleX, 0)
	fmt.Printf("%s", r.Name)
	var location string
	if ct.Name != "" {
		location = fmt.Sprintf("In %s, %s on %s going %v MPH", ct.Name, cn.State.Name, cr.Name, data.CurrentCharacter.Car.MPH)
	} else {
		location = fmt.Sprintf("In %s on %s going %v MPH", cn.State.Name, cr.Name, data.CurrentCharacter.Car.MPH)
	}
	odometer := fmt.Sprintf("Odometer: %.2f", float64(data.CurrentCharacter.Car.Odometer))
	pad := s.Width - len(location) - len(odometer)
	moveTo(0,24)
	fmt.Printf("%s%*s%s", location, pad, "", odometer)
}
