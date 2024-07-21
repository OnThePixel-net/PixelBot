const { Schema, model } = require("mongoose");

let starboard = new Schema({
  Guild: String,
  Channel: String,
  Count: Number,
});

module.exports = model("starboard", starboard);
