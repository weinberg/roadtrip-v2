import { Context } from '../util/context';
import { Character } from './character';
import { logInfo, throwError } from '../util/logging';
import { Route } from './route';

/*
model car {
  id            String     @id @default(dbgenerated("gen_random_uuid()")) @db.Uuid
  name          String?
  plate         String?    @unique
  owner_id      String?    @db.Uuid
  mph           Int?       @default(0)
  route_id      String?    @db.Uuid
  route_index   Int?       @default(0)
  node_id       String?    @db.Uuid
  node_position Int?       @default(0)
  created_at    DateTime   @default(now()) @db.Timestamp(6)
  updated_at    DateTime   @default(now()) @db.Timestamp(6)
  node          node?      @relation(fields: [node_id], references: [id], onDelete: NoAction, onUpdate: NoAction)
  character     character? @relation(fields: [owner_id], references: [id], onDelete: Cascade, onUpdate: NoAction)
  route         route?     @relation(fields: [route_id], references: [id], onDelete: NoAction, onUpdate: NoAction)
}

  type Car {
    id: String!
    name: String
    plate: String!
    riders: [Character]!
    owner: Character!
    mph: Int!
    route: Route
  }

*/

export interface Car {
  id: string;
  name: string;
  plate: string;
  riders?: Character[];
  owner?: Character;
  mph: number;
  route?: Route;
  route_id: string;
}

const owner = async (parent, args, { db }: Context): Promise<Character> => {
  let results;
  try {
    results = await db.character.findMany({
      where: {
        car: {
          some: {
            id: parent.id,
          },
        },
      },
    });
  } catch (e) {
    throwError(e as Error);
  }
  if (results?.length !== 1) {
    throwError(new Error(`Expected one character owner of car, found ${results?.length}`));
  }
  return results[0];
};

const route = async (parent: Car, args, { db }: Context) => {
  let results;
  try {
    results = await db.route.findMany({
      where: {
        id: parent.route_id,
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  if (results?.length === 0) {
    return null;
  }

  if (results?.length !== 1) {
    throwError(new Error(`Expected a single route, found ${results.length}`));
  }

  return results[0];
};

const location = async (parent: Car, args, { rtgrpc }: Context) => {
  const location = await new Promise((resolve, reject) => {
    rtgrpc.getCarLocation({ car_id: parent.id }, (error, location) => {
      if (error) {
        reject(error);
      }
      resolve(location);
    });
  });
  logInfo(`location: ${JSON.stringify(location, null, 2)}`);
  // @ts-ignore
  return { routeId: location.route_id, index: location.index, miles: location.miles };
};

export default { owner, route, location };
