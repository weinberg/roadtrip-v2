package gameServer

type (
  Node struct {
    Id string
    Miles int
  }

  Route struct {
    Id string
    Nodes []Node
  }
)
