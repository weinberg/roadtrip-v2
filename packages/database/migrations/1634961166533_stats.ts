exports.shorthands = undefined;

/**
 * Up Migration
 */
exports.up = (pgm) => {
  pgm.addColumns("car", {
    odometer: {
      type: "double precision",
      default: 0,
    },
    tripometer: {
      type: "double precision",
      default: 0,
    },
  });
};

exports.down = (pgm) => {
  pgm.dropColumns("car", ["odometer", "tripometer"]);
};
