import { Context } from '../util/context';
import { throwError } from '../util/logging';
import { Car } from './car';

export interface Character {
  id: string;
  name: string;
  created_at: Date;
  updated_at: Date;
  car?: Car;
}

const currentCharacter = async (parent, args, context): Promise<Character> => {
  return context.currentCharacter;
};

const car = async (parent, args, { currentCharacter, db }: Context) => {
  if (parent?.car) {
    return parent?.car;
  }
  const car = await db.car.findMany({
    where: {
      owner: {
        id: currentCharacter.id,
      },
    },
  });

  if (car?.length !== 1) {
    throwError(new Error(`character.car() expected one car, got ${car.length}`));
  }

  return car[0];
};

export default { car, currentCharacter };
