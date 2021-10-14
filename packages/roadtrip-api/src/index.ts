import { PrismaClient } from '@prisma/client';
import { ApolloServer } from 'apollo-server-express';
import express, { Request } from 'express';
import http from 'http';
import { logError, logInfo } from './util/logging';
import { Context, getContext } from './util/context';
import { typeDefs } from './schema';
import characterResolver from './resolvers/character';
import carResolver from './resolvers/car';
import routeResolver from './resolvers/route';
import nodeResolver from './resolvers/node';
import wayResolver from './resolvers/way';
import mapResolver from './resolvers/map';
import featureResolver from './resolvers/feature';
import { FeatureTypeEnum } from './resolvers/feature';

async function main() {
  try {
    /**
     * Database connection
     */
    const dbHost = process.env.DATABASE_HOST;
    const dbPort = process.env.DATABASE_PORT;
    const dbUser = process.env.DATABASE_USER;
    const dbPass = process.env.DATABASE_PASS;
    const dbName = process.env.DATABASE_NAME;
    const dbURL = `postgresql://${dbUser}:${dbPass}@${dbHost}:${dbPort}/${dbName}?schema=public`;

    /**
     * Prisma
     */
    const db = new PrismaClient({
      datasources: {
        db: {
          url: dbURL,
        },
      },
      // log: ['query', 'info', 'warn'],
      log: ['info', 'warn'],
    });
    // @ts-ignore
    // db.$on('query', logSql);

    /**
     * Apollo
     */

    const resolvers = {
      Query: {
        currentCharacter: characterResolver.currentCharacter,
        maps: mapResolver.maps,
      },
      Character: {
        car: characterResolver.car,
      },
      Car: {
        owner: carResolver.owner,
        route: carResolver.route,
      },
      Map: {
        routes: mapResolver.routes,
      },
      Route: {
        miles: routeResolver.miles,
        ways: routeResolver.ways,
        map: routeResolver.map,
      },
      Way: {
        nodes: wayResolver.nodes,
      },
      Node: {
        features: nodeResolver.features,
        state: nodeResolver.state,
      },
      Feature: {
        __resolveType(obj) {
          if (obj.data.type === FeatureTypeEnum.ROAD) {
            return 'Road';
          } else if (obj.data.type === FeatureTypeEnum.TOWN) {
            return 'Town';
          }
        },
      },
      Road: {
        name: featureResolver.name,
      },
      Town: {
        name: featureResolver.name,
      },
    };

    // Required logic for integrating with Express
    const app = express();
    const httpServer = http.createServer(app);
    const server = new ApolloServer({
      typeDefs,
      resolvers,
      context: async ({ req, connection }: { req: Request; connection: { context: Context } }) => {
        if (connection) {
          return connection.context;
        }

        const token: string = req?.headers?.authorization;

        if (!token) {
          return {};
        }

        let context;
        try {
          context = await getContext(db, token);
        } catch (e) {
          return {};
        }
        return context;
      },
    });

    // More required logic for integrating with Express
    await server.start();
    server.applyMiddleware({ app, path: '/' });

    const port = process.env.PORT ? process.env.PORT : 8080;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    await new Promise((resolve: any) => httpServer.listen({ port }, resolve));
    logInfo(`Server ready at http://localhost:${port}${server.graphqlPath}`);
  } catch (e) {
    logError(`Error starting application. ${e.toString()}, ${(e as Error).stack}`);
  }
}

main().catch(logError);
