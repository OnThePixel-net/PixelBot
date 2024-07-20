const { SlashCommandBuilder } = require("@discordjs/builders");
const { PermissionsBitField } = require("discord.js");

module.exports = {
  data: new SlashCommandBuilder()
    .setName("delete")
    .setDescription("Deletes a certain amount of messages from a specific user")
    .addIntegerOption((option) =>
      option
        .setName("amount")
        .setDescription("Number of messages to delete")
        .setRequired(true)
        .setMinValue(1)
        .setMaxValue(100)
    )
    .addUserOption((option) =>
      option
        .setName("user")
        .setDescription("User to delete messages from")
        .setRequired(false)
    ),

  async execute(interaction) {
    if (
      !interaction.member.permissions.has(
        PermissionsBitField.Flags.ManageMessages
      )
    ) {
      await interaction.reply({
        content: "You do not have permission to use this command.",
        ephemeral: true,
      });
      return;
    }

    const amount = interaction.options.getInteger("amount");
    const user = interaction.options.getUser("user");

    let userToDelete = null;
    if (user) {
      userToDelete = user;
    }

    if (amount > 0) {
      interaction.channel.messages
        .fetch()
        .then((messages) => {
          const now = Date.now();
          const fourteenDaysAgo = now - 14 * 24 * 60 * 60 * 1000;
          const filteredMessages = messages
            .filter(
              (message) =>
                message.createdTimestamp >= fourteenDaysAgo &&
                (userToDelete ? message.author.id === userToDelete.id : true)
            )
            .first(amount);

          if (fourteenDaysAgo) {
            interaction.reply({
              content: "❌ Messages are older than 14 days.",
              ephemeral: true,
            });
            return;
          }

          return interaction.channel.bulkDelete(filteredMessages);
        })
        .then((deletedMessages) => {
          if (deletedMessages) {
            interaction.reply({
              content: `✅ Done! Deleted ${deletedMessages.size} messages.`,
              ephemeral: true,
            });
          }
        });
    }
  },
};
