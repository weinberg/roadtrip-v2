import { Context } from '../util/context';
import { logInfo } from '../util/logging';

export interface Update {
  mph: number;
  index: number;
  miles: number;
  odometer: number;
  tripometer: number;
}

const update = async (root, args, { rtgrpc }: Context): Promise<Update> => {
  const update: Update = await new Promise((resolve, reject) => {
    rtgrpc.getUpdate({ car_id: args.id }, (error, car) => {
      if (error) {
        reject(error);
      }
      resolve(car);
    });
  });
  logInfo(`car: ${JSON.stringify(update, null, 2)}`);
  return update;
};

export default { update };
