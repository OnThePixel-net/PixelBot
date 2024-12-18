const { Client } = require("discord.js");
const mongoose = require("mongoose");
const mongoURL = process.env.MONGO_URL;
const chalk = require("chalk");

module.exports = {
  name: "ready",
  once: true,

  async execute(client) {
    if (!mongoURL) return;

    try {
      await mongoose.connect(mongoURL);
      console.log(chalk.green(`✔️ | Connected to the Database!`));
    } catch (error) {
      console.log(
        chalk.red(`✗ | Connection to Database failed! Error: ${error}`)
      );
    }

    console.log(
      chalk.green(
        `${chalk.blueBright(client.user.tag)} is logged in and online.`
      )
    );

    const activities = [
      "on OnThePixel.net",
      "dc.onthepixel.net",
      "Join now - Don't play alone!",
    ];

    setInterval(() => {
      const status = activities[Math.floor(Math.random() * activities.length)];
      client.user.setPresence({ activities: [{ name: `${status}` }] });
    }, 5000);
  },
};
