/*
model route {
  id     String         @id @default(dbgenerated("gen_random_uuid()")) @db.Uuid
  name   String?
  map_id String?        @db.Uuid
  map    map?           @relation(fields: [map_id], references: [id], onDelete: Cascade, onUpdate: NoAction)
  car    car[]
  ways   waysOnRoutes[]
}
 */

import { Context } from '../util/context';
import { throwError } from '../util/logging';

export interface Route {
  id: string;
  name: string;
  // map: Map;
  // ways: Way[];
}

const miles = () => {
  // todo fix
  return 1000;
};

const ways = async (parent: Route, args, { db }: Context) => {
  let results;
  try {
    results = await db.route.findUnique({
      where: {
        id: parent.id,
      },
      include: {
        ways: {
          include: {
            way: true,
          },
        },
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  if (!results?.ways) {
    return [];
  }

  return results?.ways.sort((a, b) => a.sequence - b.sequence).map((w) => w.way);
};

const map = async (parent, args, { db }: Context) => {
  let result;
  try {
    result = await db.route.findUnique({
      where: {
        id: parent.id,
      },
      include: {
        map: true,
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  return result?.map;
};

export default { miles, ways, map };
