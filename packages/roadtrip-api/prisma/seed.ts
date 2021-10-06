import { Prisma, PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

async function main() {
  const map = await prisma.map.create({
    data: {
      w: 80,
      h: 25,
      glyphs: String.raw`                                                                                    ,__                                                     _,                       \~\|~~~~---___              ,                          | \                       |            ~~~~~~~~~~~~~- ~~---,                 __/   >                     /~                                \`~\_             /~    ,'                     |                                     / /~)    __/      \,                     /                                     | | '~\  |        ,-'                     |                                     | |   /_-'       ~                        |                                     \`-'             /                         |                                                    |\`                         ',                                                   |                           |                                                   \                           ',                                                  /                            '_                                                /                               \                                             /~                                 ~~~-                                        /                                       '-,_                                    \                                           \`~'~~~\                   ,~~~~-~~,  \                                                 \/~\      /~~~\`---\`         |  \                                                    \    /                   \  |                                                    \  |                     '\'                                                     \`~'                                                                                                                               `,
    },
  });

  const wayNodes = [
    {
      name: 'I-90',
      nodes: [
        [7, 3, 10, '═'],
        [8, 3, 10, '╗'],
      ],
    },
    {
      name: 'I-82',
      nodes: [
        [8, 4, 10, '╚'],
        [9, 4, 10, '╗'],
        [9, 5, 10, '╚'],
        [10, 5, 10, '╗'],
        [10, 6, 10, '╚'],
      ],
    },
    {
      name: 'I-84',
      nodes: [
        [11, 6, 10, '╗'],
        [11, 7, 10, '╚'],
        [12, 7, 10, '═'],
        [13, 7, 10, '╗'],
        [13, 8, 10, '╚'],
        [14, 8, 10, '═'],
        [15, 8, 10, '═'],
        [16, 8, 10, '╗'],
        [16, 9, 10, '╚'],
      ],
    },
    {
      name: 'I-80',
      nodes: [
        [17, 9, 10, '═'],
        [18, 9, 10, '═'],
        [19, 9, 10, '═'],
        [20, 9, 10, '╝'],
        [20, 8, 10, '╔'],
        [21, 8, 10, '═'],
        [22, 8, 10, '═'],
        [23, 8, 10, '═'],
      ],
    },
    {
      name: 'I-25',
      nodes: [
        [24, 8, 10, '╗'],
        [24, 9, 10, '║'],
        [24, 10, 10, '║'],
      ],
    },
  ];

  const route = await prisma.route.create({
    data: {
      map_id: map.id,
      name: 'Seattle, Washington to Denver, Colorado',
      ways: {
        create: wayNodes.map((wayNode, i) => ({
          sequence: i,
          way: {
            create: {
              name: wayNode.name,
              nodes: {
                create: wayNode.nodes.map((n, j) => ({
                  sequence: j,
                  node: {
                    create: {
                      x: n[0] as number,
                      y: n[1] as number,
                      miles: n[2] as number,
                      features: { create: [{ glyph: n[3], data: { type: 'ROAD' } as Prisma.JsonObject }] },
                    },
                  },
                })),
              },
            },
          },
          /*
          nodes: {
            create: wayNode.nodes.map((n, j) => ({
              sequence: j,
              node: {
                create: {
                  x: n[0] as number,
                  y: n[1] as number,
                  miles: n[2] as number,
                  features: { create: [{ glyph: n[3], data: { type: 'ROAD' } as Prisma.JsonObject }] },
                },
              },
            })),
          },

           */
        })),
      },
    },
  });

  /*
  for (const wn of wayNodes) {
    await prisma.way.create({
      data: {
        name: wn.name,
        nodes: {
          create: wn.nodes.map((n, i) => ({
            sequence: i,
            node: {
              create: {
                x: n[0] as number,
                y: n[1] as number,
                miles: n[2] as number,
                features: { create: [{ glyph: n[3], data: { type: 'ROAD' } as Prisma.JsonObject }] },
              },
            },
          })),
        },
      },
    });
  }
 */
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });
