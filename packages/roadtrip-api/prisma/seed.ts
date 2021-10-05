import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

async function main() {
  const map = await prisma.map.create({
    data: {
      w: 80,
      h: 25,
      glyphs: String.raw`                                                                                    ,__                                                     _,                       \~\|~~~~---___              ,                          | \                       |            ~~~~~~~~~~~~~- ~~---,                 __/   >                     /~                                \`~\_             /~    ,'                     |                                     / /~)    __/      \,                     /                                     | | '~\  |        ,-'                     |                                     | |   /_-'       ~                        |                                     \`-'             /                         |                                                    |\`                         ',                                                   |                           |                                                   \                           ',                                                  /                            '_                                                /                               \                                             /~                                 ~~~-                                        /                                       '-,_                                    \                                           \`~'~~~\                   ,~~~~-~~,  \                                                 \/~\      /~~~\`---\`         |  \                                                    \    /                   \  |                                                    \  |                     '\'                                                     \`~'                                                                                                                               `,
    },
  });
  const stdRoute = await prisma.route.create({
    data: {
      map_id: map.id,
      name: 'Seattle, Washington to Denver, Colorado',
      ways: {
        create: [
          {
            way: {
              create: {
                name: 'I-90',
                nodes: {
                  create: [
                    {
                      node: {
                        create: {
                          x: 6,
                          y: 2,
                          miles: 10,
                        },
                      },
                    },
                  ],
                },
              },
            },
          },
        ],
      },
    },
  });

  const seattleToDenverNodes = [{ x: 6, y: 2, c: '•' }];

  console.log({ map, stdRoute });
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });
