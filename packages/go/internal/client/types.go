package client

import "github.com/hasura/go-graphql-client"

type (
	State struct {
		Abbreviation graphql.String
		Name         graphql.String
	}

	Road struct {
    Typename string `graphql:"__typename"`
		Name       graphql.String
		Glyph      graphql.String
	}
	Town struct {
		Typename string `graphql:"__typename"`
		Name     graphql.String
		Glyph    graphql.String
	}

	Node struct {
		X        int
		Y        int
		Tz       int
		State    State
		Features []struct {
			Road `graphql:"... on Road"`
			Town `graphql:"... on Town"`
		}
	}

	Way struct {
		Name  graphql.String
		Nodes []Node
	}

	Map struct {
		W     int
		H     int
		Image graphql.String
	}

	Route struct {
		Name graphql.String
		Map  Map
		Ways []Way
	}

	Car struct {
		Id         graphql.ID
		Name       graphql.String
		Route      Route
		Location   Location
		MPH        graphql.Float
		Odometer   graphql.Float
		Tripometer graphql.Float
	}

	Location struct {
		RouteId graphql.String
		Index   int
		Miles   graphql.Float
	}

	CurrentCharacter struct {
		Name graphql.String
		Car  Car
	}

	Update struct {
		MPH        graphql.Float
		Miles      graphql.Float
		Index      int
		Odometer   graphql.Float
		Tripometer graphql.Float
	}
)
