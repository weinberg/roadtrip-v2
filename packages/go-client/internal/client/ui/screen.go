package ui

import (
  "fmt"
  "github.com/brickshot/roadtrip-v2/go-client/internal/client"
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
  CurrentCharacter types.CurrentCharacter
}

func Render(data RenderData) {
  cls()
  moveTo(0,0)
  m := data.CurrentCharacter.Car.Route.Map
  Outer:
    for r := 0; r < m.H; r++ {
      for c := 0; c < m.W; c++ {
        i :=r*m.W+c;
        if i >= len(m.Image) {
          break Outer
        }
        fmt.Printf("%c", m.Image[i])
      }
      fmt.Printf("\n")
    }

  r := data.CurrentCharacter.Car.Route
  moveTo(0,0)
  fmt.Printf("Route: %s", r.Name)

  for _, w := range r.Ways {
    for _, n := range w.Nodes {
      for _, f := range n.Features {
        if f.Road.Glyph != "" {
          moveTo(n.X, n.Y)
          fmt.Printf("%v", f.Road.Glyph)
        }
      }
    }
  }
  moveTo(7,0)
  fmt.Printf("x")
  moveTo(0,3)
  fmt.Printf("y")
  moveTo(0,25)
}

