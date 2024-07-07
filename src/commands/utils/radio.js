const { SlashCommandBuilder } = require("discord.js");
const {
  joinVoiceChannel,
  createAudioPlayer,
  createAudioResource,
  NoSubscriberBehavior,
  StreamType,
} = require("@discordjs/voice");

module.exports = {
  data: new SlashCommandBuilder()
    .setName("radio")
    .setDescription("Plays OnThePixel Radio"),
  async execute(interaction) {
    if (!interaction.member.voice.channel) {
      return interaction.reply(
        "You need to be in a voice channel to use this command."
      );
    }

    const connection = joinVoiceChannel({
      channelId: interaction.member.voice.channel.id,
      guildId: interaction.guild.id,
      adapterCreator: interaction.guild.voiceAdapterCreator,
    });

    const player = createAudioPlayer({
      behaviors: {
        noSubscriber: NoSubscriberBehavior.Pause,
      },
    });

    // Improved handling of audio resources
    const resource = createAudioResource("https://stream.laut.fm/onthepixel", {
      inputType: StreamType.Arbitrary,
    });

    player.play(resource);
    connection.subscribe(player);

    player.on("error", (error) => {
      console.error("Error from the audio player:", error);
      interaction.followUp("There was an error playing the radio.");
    });

    await interaction.reply("Now playing OnThePixel Radio!");

    // Handle disconnection
    player.on("stateChange", (oldState, newState) => {
      if (newState.status === "idle") {
        connection.destroy();
      }
    });
  },
};
