const {
  Events,
  ActionRowBuilder,
  ButtonBuilder,
  ButtonStyle,
  ComponentType,
  EmbedBuilder,
} = require("discord.js");

module.exports = {
  name: Events.InteractionCreate,
  async execute(interaction, client) {
    if (!interaction.commandName) return;

    var sendGuild = await client.guilds.fetch(process.env.GUILD_ID);
    var sendChannel = await sendGuild.channels.fetch(process.env.LOG_CHANNEL);

    var command = interaction.commandName;
    var user = interaction.user;
    var channel = interaction.channel;

    const embed = new EmbedBuilder()
      .setColor("Green")
      .setTitle(`Command Used`)
      .setDescription("An interaction command has been used.")
      .addFields({ name: "Command", value: `\`${command}\`` })
      .addFields({
        name: "Channel:",
        value: `\`${channel.name}\` (${channel.id})`,
      })
      .addFields({
        name: "User:",
        value: `\`${user.username}\` (${user.id})`,
      })
      .setFooter({ text: "Interaction Use Logger" })
      .setTimestamp();

    var msg = await sendChannel.send({
      embeds: [embed],
    });
  },
};
