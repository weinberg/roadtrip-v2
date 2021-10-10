import { Car } from './car';
import { Feature } from './feature';
import { Context } from '../util/context';
import { throwError } from '../util/logging';

/*
model node {
  id         String        @id @default(dbgenerated("gen_random_uuid()")) @db.Uuid
  x          Int?          @default(0)
  y          Int?          @default(0)
  miles      Int?          @default(0)
  created_at DateTime      @default(now()) @db.Timestamp(6)
  updated_at DateTime      @default(now()) @db.Timestamp(6)
  cars       car[]
  features   feature[]
  ways       nodesOnWays[]
}
*/

export interface Node {
  id: string;
  x: number;
  y: number;
  miles: number;
  cars: Car[];
  features: Feature[];
}

const features = async (parent, args, { db }: Context) => {
  let results;
  try {
    results = await db.node.findUnique({
      where: {
        id: parent.id,
      },
      include: {
        features: true,
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  return results?.features || [];
};

const state = async (parent, args, { db }: Context) => {
  if (parent?.state) {
    return parent.state;
  }

  let results;
  try {
    results = await db.node.findUnique({
      where: {
        id: parent.id,
      },
      include: {
        state: true,
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  return results?.state;
};

export default { features, state };
