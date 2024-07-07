const { SlashCommandBuilder } = require("@discordjs/builders");
const { PermissionsBitField } = require("discord.js");

module.exports = {
  data: new SlashCommandBuilder()
    .setName("edit")
    .setDescription("Edits a message sent by the bot.")
    .addStringOption((option) =>
      option
        .setName("message_id")
        .setDescription("The ID of the message to edit")
        .setRequired(true)
    )
    .addStringOption((option) =>
      option
        .setName("new_text")
        .setDescription("The new text for the message")
        .setRequired(true)
    ),
  async execute(interaction, client) {
    if (
      !interaction.member.permissions.has(
        PermissionsBitField.Flags.ManageMessages
      )
    )
      return await interaction.reply({
        content: "You need the Manage Messages permission to use this command.",
        ephemeral: true,
      });

    const messageId = interaction.options.getString("message_id");
    const newText = interaction.options.getString("new_text");
    const channel = interaction.channel;

    try {
      const message = await channel.messages.fetch(messageId);
      if (message.author.id !== client.user.id) {
        return await interaction.reply({
          content: "I can only edit messages that were sent by me.",
          ephemeral: true,
        });
      }

      await message.edit(newText);
      await interaction.reply({
        content: "Message edited successfully!",
        ephemeral: true,
      });
    } catch (error) {
      console.error(error);
      await interaction.reply({
        content: "An error occurred while trying to edit the message.",
        ephemeral: true,
      });
    }
  },
};
