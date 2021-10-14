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
  m := data.CurrentCharacter.Car.Route.Map
  for r := 0; r < m.H-1; r++ {
    for c := 0; c < m.W-1; c++ {
      fmt.Printf("%c", m.Image[r*m.W+c])
    }
    fmt.Printf("\n")
  }
}

