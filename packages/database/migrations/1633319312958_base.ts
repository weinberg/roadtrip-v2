/* eslint-disable camelcase */

exports.shorthands = undefined;

/**
 * Up Migration
 */
exports.up = (pgm) => {
  /******************************
   * Entities
   ******************************/

  /**
   * Map
   */
  pgm.createTable("map", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    name: {
      type: "string",
    },
    w: {
      type: "int",
    },
    h: {
      type: "int",
    },
    image: {
      type: "string",
    },
  });

  /**
   * State
   */
  pgm.createTable("state", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    name: {
      type: "string",
    },
    abbreviation: {
      type: "string",
    },
  });

  /**
   * Nodes
   */
  pgm.createTable("node", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    x: {
      type: "int",
      default: 0,
    },
    y: {
      type: "int",
      default: 0,
    },
    state_id: {
      type: "uuid",
      references: '"state"',
      onDelete: "set null",
    },
    tz: {
      type: "int",
      default: -8,
    },
    miles: {
      type: "int",
      default: 0,
    },
    created_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
    updated_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
  });

  /**
   * Feature
   */
  pgm.createTable("feature", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    node_id: {
      type: "uuid",
      notNull: true,
      references: '"node"',
      onDelete: "cascade",
    },
    data: {
      type: "jsonb",
    },
    glyph: {
      type: "string",
      default: " ",
    },
    created_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
    updated_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
  });

  /**
   * Way
   */
  pgm.createTable("way", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    name: {
      type: "string",
    },
  });

  /**
   * NodesOnWay join table
   */
  pgm.createTable("nodesOnWays", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    node_id: {
      type: "uuid",
      references: '"node"',
    },
    way_id: {
      type: "uuid",
      references: '"way"',
    },
    sequence: {
      type: "int",
      default: 0,
    },
  });

  /**
   * Route
   */
  pgm.createTable("route", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    name: {
      type: "string",
    },
    map_id: {
      type: "uuid",
      references: '"map"',
      onDelete: "cascade",
    },
  });

  /**
   * WaysOnRoute join table
   */
  pgm.createTable("waysOnRoutes", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    way_id: {
      type: "uuid",
      references: '"way"',
    },
    route_id: {
      type: "uuid",
      references: '"route"',
    },
    sequence: {
      type: "int",
      default: 0,
    },
  });

  /**
   * Character
   */
  pgm.createTable("character", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    token: {
      type: "uuid",
      unique: true,
      notNull: true,
    },
    name: { type: "string" },
    created_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
    updated_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
  });

  /**
   * Car
   */
  pgm.createTable("car", {
    id: {
      type: "uuid",
      primaryKey: true,
      default: pgm.func("gen_random_uuid()"),
      unique: true,
    },
    name: {
      type: "string",
    },
    plate: {
      type: "string",
      unique: true,
    },
    owner_id: {
      type: "uuid",
      references: '"character"',
      onDelete: "cascade",
    },
    mph: {
      type: "int",
      default: 0,
    },
    // route this car is currently on
    route_id: {
      type: "uuid",
      references: '"route"',
    },
    // index of current node if all route.way[].nodes[] were laid out sequentially
    route_index: {
      type: "int",
      default: 0,
    },
    // current node
    node_id: {
      type: "uuid",
      references: '"node"',
    },
    // car position inside the current node
    node_position: {
      type: "int",
      default: 0,
    },
    created_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
    updated_at: {
      type: "timestamp",
      notNull: true,
      default: pgm.func("current_timestamp"),
    },
  });
};

/**
 * Down migration
 */
exports.down = (pgm) => {
  pgm.dropTable("map", { cascade: true, ifExists: true });
  pgm.dropTable("node", { cascade: true, ifExists: true });
  pgm.dropTable("feature", { cascade: true, ifExists: true });
  pgm.dropTable("way", { cascade: true, ifExists: true });
  pgm.dropTable("nodesOnWays", { cascade: true, ifExists: true });
  pgm.dropTable("route", { cascade: true, ifExists: true });
  pgm.dropTable("waysOnRoutes", { cascade: true, ifExists: true });
  pgm.dropTable("character", { cascade: true, ifExists: true });
  pgm.dropTable("car", { cascade: true, ifExists: true });
};
