import { Context } from '../util/context';
import { Character } from './character';
import { logInfo, throwError } from '../util/logging';
import { Route } from './route';

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

const car = async (root, args, { rtgrpc }: Context) => {
  const car = await new Promise((resolve, reject) => {
    rtgrpc.getCar({ car_id: args.id }, (error, car) => {
      if (error) {
        reject(error);
      }
      resolve(car);
    });
  });
  logInfo(`car: ${JSON.stringify(car, null, 2)}`);
  return {
    id: args.id,
    // @ts-ignore
    odometer: car.odometer,
    // @ts-ignore
    tripometer: car.tripometer,
    // @ts-ignore
    location: { routeId: car.location.route_id, index: car.location.index, miles: car.location.miles },
    // @ts-ignore
    mph: car.mph,
  };
};

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

export default { car, owner, route };
