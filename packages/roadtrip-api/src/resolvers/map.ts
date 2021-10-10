/*
model map {
  id     String  @id @default(dbgenerated("gen_random_uuid()")) @db.Uuid
  name   String?
  w      Int?
  h      Int?
  image String?
  routes  route[]
}

 */

import { Context } from '../util/context';
import { throwError } from '../util/logging';

export interface Map {
  id: string;
  name: string;
  w: number;
  h: number;
  image: string;
  // routes: Route[]
}

const maps = async (parent, args, { db }: Context) => {
  let results;
  try {
    results = await db.map.findMany();
  } catch (e) {
    throwError(e as Error);
  }

  return results;
};

const routes = async (parent, args, { db }: Context) => {
  let results;
  try {
    results = await db.route.findMany({
      where: {
        map_id: parent.id,
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  return results;
};

export default { maps, routes };
