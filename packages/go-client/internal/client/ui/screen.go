package ui

import (
  "fmt"
)

var esc = "\033"

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
