package gameServer

type (
  Node struct {
    Id string
    Miles float64
  }

  Route struct {
    Id string
    Nodes []Node
  }

  Car struct {
    Id string
    RouteId string
    RouteIndex int32
    NodeMiles float64
    Mph float64
  }
)
