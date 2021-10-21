import { gql } from 'apollo-server-express';

const typeDefs = gql`
  #
  # Types
  #

  type Character {
    id: String!
    "Character's token"
    token: String!
    "Character's name"
    name: String!
    "Character's Car"
    car: Car!
  }

  type Location {
    routeId: String!
    "Index into the route nodes if all route->ways->nodes were in an array."
    index: Int!
    "Miles into the current node"
    miles: Float!
  }

  type Car {
    id: String!
    "Car name"
    name: String
    plate: String!
    "All characters riding in the car"
    riders: [Character]!
    "The owner of the Car"
    owner: Character!
    "Velocity in MPH"
    mph: Int!
    "Current route"
    route: Route
    "Current Location"
    location: Location
  }

  type State {
    id: String!
    "The name of the State"
    name: String!
    "The abbreviation for the state"
    abbreviation: String
  }

  type Town {
    id: String!
    "Name of town"
    name: String
    "UTF8 character glyph to represent this node"
    glyph: String!
  }

  type Road {
    id: String!
    "Name of road"
    name: String
    "UTF8 character glyph to represent this node"
    glyph: String!
  }

  union Feature = Town | Road

  # A node is a single character position on the map
  type Node {
    "0 indexed x position"
    x: Int!
    "0 indexed y position"
    y: Int!
    "State this node is in"
    state: State!
    "Timezone offset from UTC"
    tz: Int!
    "Features at this node"
    features: [Feature!]!
    "Length in miles of this node"
    miles: Float!
  }

  # A way is like a route segment - so like a highway
  type Way {
    id: String!
    "Nodes in this way"
    nodes: [Node!]!
    "Name of the way - typically the name of the highway this way represents"
    name: String!
    "Length of this way. This is equal to the sum of the miles of all the nodes in the way."
    miles: Float!
  }

  # A route is an ordered collection of ways.
  # The name is a description of the route like "Seattle to Denver".
  type Route {
    id: String!
    name: String!
    ways: [Way!]!
    "Length of this route. This is equivalent to the sum of the miles of all the ways in the route."
    miles: Float!
    "The map this route is on."
    map: Map!
  }

  type Map {
    id: String!
    """
    String of all the glyphs representing a background image of the map.
    The string is w*h in length with no carriage returns or separators between rows.
    """
    image: String!
    "Width of map in characters"
    w: Int!
    "Height of map in characters"
    h: Int!
    "All the routes in the map."
    routes: [Route!]!
  }

  #
  # Query
  #
  type Query {
    currentCharacter: Character
    character(id: ID!): Character
    maps: [Map!]!
    map(id: ID): Map
    routes: [Route!]!
    route(id: ID!): Route
  }

  #
  # Mutation
  #
  type Mutation {
    createCharacter(id: ID!, name: String!): Character
  }
`;

export { typeDefs };
