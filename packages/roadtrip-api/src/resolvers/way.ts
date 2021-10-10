/*
model way {
  id     String         @id @default(dbgenerated("gen_random_uuid()")) @db.Uuid
  name   String?
    nodes  nodesOnWays[]
  routes waysOnRoutes[]
}
*/

import { throwError } from '../util/logging';
import { Context } from '../util/context';

export interface Way {
  id: string;
  name: string;
  // nodes: Node[];
}

const nodes = async (parent, args, { db }: Context) => {
  let results;
  try {
    results = await db.way.findUnique({
      where: {
        id: parent.id,
      },
      include: {
        nodes: {
          include: {
            node: true,
          },
        },
      },
    });
  } catch (e) {
    throwError(e as Error);
  }

  if (!results?.nodes?.length) {
    return [];
  }

  return results.nodes.sort((a, b) => a.sequence - b.sequence).map((n) => n.node);
};

export default { nodes };
