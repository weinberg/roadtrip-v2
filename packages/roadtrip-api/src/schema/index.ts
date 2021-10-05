import { gql } from 'apollo-server';

const typeDefs = gql`
  #
  # Types
  #

  type Character {
    id: String!
    "Character's name"
    name: String!
    "Character's Car"
    car: Car!
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
  }

  type State {
    id: String!
    "The name of the State"
    name: String!
  }

  type City {
    id: String!
    "Name of city"
    name: String!
    "UTF8 character glyph to represent this node"
    c: String!
  }

  type Road {
    id: String!
    "Name of road"
    name: String!
    "UTF8 character glyph to represent this node"
    c: String!
  }

  union Feature = City | Road

  # A node is a single character position on the map
  type Node {
    "0 indexed x position"
    x: Int!
    "0 indexed y position"
    y: Int!
    "State this node is in"
    state: State!
    "Features at this node"
    feature: [Feature!]!
    "Length in miles of this node"
    miles: Int!
  }

  # A way is like a route segment - so like a highway
  type Way {
    id: String!
    "Nodes in this way"
    nodes: [Node!]!
    "Name of the way - typically the name of the highway this way represents"
    name: String!
    "Length of this way. This is equal to the sum of the miles of all the nodes in the way."
    miles: Int!
  }

  # A route is an ordered collection of ways.
  # The name is a description of the route like "Seattle to Denver".
  type Route {
    id: String!
    name: String!
    ways: [Way!]!
    "Length of this route. This is equivalent to the sum of the miles of all the ways in the route."
    miles: Int!
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
    "All the nodes in the map."
    nodes: [Node!]!
  }

  #
  # Query
  #
  type Query {
    character(id: ID!): Character
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
