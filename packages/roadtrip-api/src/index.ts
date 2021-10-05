import { ApolloServer } from 'apollo-server';
import { typeDefs } from './schema';

// The ApolloServer constructor requires two parameters: your schema
// definition and your set of resolvers.
const server = new ApolloServer({ typeDefs: [typeDefs], resolvers });

// The `listen` method launches a web server.
server.listen().then(({ url }) => {
  // eslint-disable-next-line no-console
  console.log(`🚀  Server ready at ${url}`);
});
