import { Prisma, PrismaClient } from "@prisma/client";
import * as fs from "fs";

const prisma = new PrismaClient();

async function main() {
  // read map image from file
  let maptext = fs.readFileSync(`${__dirname}/map.txt`, "utf8");
  let rows = maptext.split(/\n/);
  let image;
  image = rows.map((r) => r.padEnd(80, " ")).join("");

  const map = await prisma.map.create({
    data: {
      w: 80,
      h: 25,
      image,
    },
  });

  // washington - I90, I82
  // oregon - I84
  // idaho - I84
  // utah - I84 I80
  // wyoming I80
  // colorado I25

  const WA = await prisma.state.create({
    data: {
      name: "Washington",
      abbreviation: "WA",
    },
  });
  const OR = await prisma.state.create({
    data: {
      name: "Oregon",
      abbreviation: "OR",
    },
  });
  const UT = await prisma.state.create({
    data: {
      name: "Utah",
      abbreviation: "UT",
    },
  });
  const ID = await prisma.state.create({
    data: {
      name: "Idaho",
      abbreviation: "ID",
    },
  });
  const WY = await prisma.state.create({
    data: {
      name: "Wyoming",
      abbreviation: "WY",
    },
  });
  const CO = await prisma.state.create({
    data: {
      name: "Colorado",
      abbreviation: "CO",
    },
  });

  const wayNodes = [
    {
      name: "I-90",
      nodes: [
        [7, 3, 10, "═", WA, -8],
        [8, 3, 10, "╗", WA, -8],
      ],
    },
    {
      name: "I-82",
      nodes: [
        [8, 4, 10, "╚", WA, -8],
        [9, 4, 10, "╗", WA, -8],
      ],
    },
    {
      name: "I-84",
      nodes: [
        [9, 5, 10, "╚", OR, -8],
        [10, 5, 10, "╗", OR, -8],
        [10, 6, 10, "╚", OR, -8],
        [11, 6, 10, "╗", OR, -8],
        [11, 7, 10, "╚", OR, -8],
        [12, 7, 10, "═", OR, -7],
        [13, 7, 10, "╗", ID, -7],
        [13, 8, 10, "╚", ID, -7],
        [14, 8, 10, "═", ID, -7],
        [15, 8, 10, "═", ID, -7],
        [16, 8, 10, "╗", UT, -7],
        [16, 9, 10, "╚", UT, -7],
      ],
    },
    {
      name: "I-80",
      nodes: [
        [17, 9, 10, "═", UT, -7],
        [18, 9, 10, "═", WY, -7],
        [19, 9, 10, "═", WY, -7],
        [20, 9, 10, "╝", WY, -7],
        [20, 8, 10, "╔", WY, -7],
        [21, 8, 10, "═", WY, -7],
        [22, 8, 10, "═", WY, -7],
        [23, 8, 10, "═", WY, -7],
      ],
    },
    {
      name: "I-25",
      nodes: [
        [24, 8, 10, "╗", CO, -7],
        [24, 9, 10, "║", CO, -7],
        [24, 10, 10, "║", CO, -7],
      ],
    },
  ];

  // @ts-ignore
  const route = await prisma.route.create({
    data: {
      map_id: map.id,
      name: "Seattle, Washington to Denver, Colorado",
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
                      tz: n[5] as number,
                      state: {
                        connect: {
                          id: (n[4] as any).id as string,
                        },
                      },
                      features: {
                        create: [
                          {
                            glyph: n[3],
                            data: {
                              type: "ROAD",
                              name: wayNode.name,
                            } as Prisma.JsonObject,
                          },
                          i === 0 && j === 0
                            ? {
                                glyph: "*",
                                data: {
                                  type: "TOWN",
                                  name: "Seattle",
                                } as Prisma.JsonObject,
                              }
                            : undefined,
                        ],
                      },
                    },
                  },
                })),
              },
            },
          },
        })),
      },
    },
    include: {
      ways: {
        include: {
          way: {
            include: {
              nodes: {
                include: {
                  way: true,
                },
              },
            },
          },
        },
      },
    },
  });

  const car = await prisma.car.create({
    data: {
      name: "Herbie",
      plate: "ABC-XYZ",
      route: {
        connect: {
          id: route.id,
        },
      },
      route_index: 0,
      node_miles: 0,
      mph: 60,
    },
  });

  await prisma.character.create({
    data: {
      name: "Josh",
      token: "49ba3caa-d0da-4e81-9ee0-47ce29d05e69",
      car: {
        connect: {
          id: car.id,
        },
      },
    },
  });
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });
