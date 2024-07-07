require("dotenv").config();
const { token } = process.env;
const GiveawaysManager = require("./giveaway");
const { Client, Collection, GatewayIntentBits } = require("discord.js");
const fs = require("fs");
const counting = require("./Schemas/countingSchema");
// const axios = require("axios");

const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
    GatewayIntentBits.GuildMembers,
    GatewayIntentBits.GuildMessageReactions,
    GatewayIntentBits.DirectMessages,
  ],
});
client.commands = new Collection();
client.commandArray = [];

const functionFolders = fs.readdirSync(`src/functions`);
for (const folder of functionFolders) {
  const functionFiles = fs
    .readdirSync(`src/functions/${folder}`)
    .filter((file) => file.endsWith(".js"));
  for (const file of functionFiles)
    require(`./functions/${folder}/${file}`)(client);
}

// Welcome
client.on("guildMemberAdd", async (member) => {
  const channelID = await process.env.channelID;
  const channel = member.guild.channels.cache.get(channelID);
  const message = `**Welcome to OnThePixel.net, ${member}!**`;

  if (channelID == null) return;

  channel.send(message);
});
// Welcome end

// Auto role
client.on("guildMemberAdd", async (member) => {
  const role = process.env.roleID;
  const giveRole = await member.guild.roles.cache.get(role);

  member.roles.add(giveRole);
});
// Auto role end

// Reaction role
const reactions = require("./Schemas/reactionRoleSchema");
client.on("messageReactionAdd", async (reaction, user) => {
  if (!reaction.message.guildId) return;
  if (user.bot) return;

  let cID = `<:${reaction.emoji.name}:${reaction.emoji.id}>`;
  if (!reaction.emoji.id) cID = reaction.emoji.name;

  const data = await reactions.findOne({
    Guild: reaction.message.guildId,
    Message: reaction.message.id,
    Emoji: cID,
  });

  if (!data) return;

  const guild = await client.guilds.cache.get(reaction.message.guildId);
  const member = await guild.members.cache.get(user.id);

  try {
    await member.roles.add(data.Role);
  } catch (e) {
    return;
  }
});

client.on("messageReactionRemove", async (reaction, user) => {
  if (!reaction.message.guildId) return;
  if (user.bot) return;

  let cID = `<:${reaction.emoji.name}:${reaction.emoji.id}>`;
  if (!reaction.emoji.id) cID = reaction.emoji.name;

  const data = await reactions.findOne({
    Guild: reaction.message.guildId,
    Message: reaction.message.id,
    Emoji: cID,
  });

  if (!data) return;

  const guild = await client.guilds.cache.get(reaction.message.guildId);
  const member = await guild.members.cache.get(user.id);

  try {
    await member.roles.remove(data.Role);
  } catch (e) {
    return;
  }
});
// Reaction role end

// Giveaway
client.giveawayManager = new GiveawaysManager(client, {
  default: {
    botsCanWin: false,
    embedColor: "#a200ff",
    embedColorEnd: "#550485",
    reaction: "🎉",
  },
});
// Giveaway end

// Counting
client.on("messageCreate", async (message) => {
  if (!message.guild) return;
  if (message.author.bot) return;

  const data = await counting.findOne({ Guild: message.guild.id });
  if (!data) return;
  else {
    if (message.channel.id !== data.channel) return;

    const number = Number(message.content);

    if (number !== data.Number) {
      return message.react(`❌`);
    } else if (data.LastUser === message.author.id) {
      message.react(`❌`);
      await message.reply(`❌ You Can Not Count More Than Once In A Row!`);
    } else {
      await message.react(`✅`);

      data.LastUser = message.author.id;
      data.Number++;
      await data.save();
    }
  }
});
// Counting end

// Response start
client.on("messageCreate", async (message) => {
  if (message.author.bot) return;

  const botMentioned = message.mentions.has(client.user.id);
  const containsPixelbot = message.content.toLowerCase().includes("pixelbot");
  const messageContent = message.content.toLowerCase();

  const triggerWords = ["hi", "guten tag", "hallo", "moin"];
  const responsePhrases = [
    "Moin moin!",
    "Hi, wie geht's?",
    "Guten Tag, Reisender!",
    "Howdy, Partner!",
  ];
  const mentionResponsePhrases = ["Hmm?", "Yeah?", "Was ist denn?"];
  if (botMentioned || containsPixelbot) {
    const response =
      mentionResponsePhrases[
        Math.floor(Math.random() * mentionResponsePhrases.length)
      ];
    message.channel.send(response);
    return;
  }

  const containsTriggerWord = triggerWords.some((word) => {
    const regex = new RegExp(`\\b${word}\\b`, "i");
    return regex.test(messageContent);
  });

  if (containsTriggerWord) {
    const response =
      responsePhrases[Math.floor(Math.random() * responsePhrases.length)];
    message.channel.send(response);
    return;
  }

  const jokes = [
    "Meine Gitarre ist zu leise um mittelalterliche Musik zu spielen. Ich brauch eine Laute.",
    "Mir ist kalt - ich geh mal in die Ecke. Da sind 90 Grad.",
    "Sagt die eine Kerze zur anderen: Du, ist Wasser eigentlich gefährlich? - Sagt die andere: Davon kannst Du ausgehen.",
  ];

  const creatorPhrases = [
    "Ich wurde von einem sehr schlauem Entwickler erschaffen.",
    "Ein kluger Köpfe hat mich zum Leben erweckt.",
    "Ich bin das Produkt vieler Stunden harter Arbeit und Kreativität von <@982984144567017493>.",
    "Es war ein langer Weg, aber <@982984144567017493> hat mich erschaffen.",
    "Ich bin das Werk von <@982984144567017493>.",
  ];

  const ipPhrases = [
    "Klar, die IP ist onthepixel.net!",
    "Du kannst dich mit der IP onthepixel.net zum Server verbinden.",
    "Du suchst die IP? onthepixel.net ist es. Wir sehen uns da!",
  ];

  const wantsJoke = /erzähl.*witz|erzähle.*witz/.test(messageContent);
  const asksForIP = /\bip\b/.test(messageContent);
  const asksForCreator =
    /wer hat dich gemacht/ |
    /wer hat dich erstellt/ |
    /wer ist dein macher/ |
    /wer hat dich geschaffen/.test(messageContent);

  if (!botMentioned && !containsPixelbot) {
    return;
  }
  if (asksForCreator) {
    const response =
      creatorPhrases[Math.floor(Math.random() * creatorPhrases.length)];
    message.channel.send(response);
    return;
  }
  if (containsTriggerWord) {
    await message.channel.sendTyping();
    const randomResponse =
      responsePhrases[Math.floor(Math.random() * responsePhrases.length)];
    await message.reply(randomResponse);
  } else if (wantsJoke) {
    await message.channel.sendTyping();
    const randomJoke = jokes[Math.floor(Math.random() * jokes.length)];
    await message.reply(randomJoke);
  } else if (asksForIP) {
    await message.channel.sendTyping();
    const randomIPPhrase =
      ipPhrases[Math.floor(Math.random() * ipPhrases.length)];
    await message.reply(randomIPPhrase);
  }
});
// Response end

// Chat Bot logic
// RAPIDAPI BARD API ISNT WORKING ATM

// client.on("messageCreate", async (message) => {
//   if (message.author.bot) return;

//   await message.channel.sendTyping();

//   let input = {
//     method: 'GET',
//     url: 'https://google-bard1.p.rapidapi.com/',
//     headers: {
//       message: message.content,
//       'X-RapidAPI-Key': '', // YOUR OWN RAPIDAPI KEY
//       'X-RapidAPI-Host': 'google-bard1.p.rapidapi.com'
//     },
//   };
//   try {
//     const output = await axios.request(input);
//     const response = output.data.response;

//     if (response.length > 2000) {
//       const chunks = response.match(/.{1,2000}/g);

//       for (let i = 0; i < chunks.length; i++) {
//         message.author.send(chunks[i]).catch((err) => {
//           console.error(err);
//           message.author
//             .send(
//               "I'm heaving trouble finding that request. Because I am an AI on Discord, I don't have time to process long requests."
//             )
//             .catch((err) => {
//               console.error(err);
//             });
//         });
//       }
//     } else {
//       await message.author.send(response).catch((err) => {
//         console.error(err);
//         message.author
//           .send(
//             "I'm heaving trouble finding that request. Because I am an AI on Discord, I don't have time to process long requests."
//           )
//           .catch((err) => {
//             console.error(err);
//           });
//       });
//     }
//   } catch (e) {
//     console.error(e);
//     message.author
//       .send(
//         "I'm heaving trouble finding that request. Because I am an AI on Discord, I don't have time to process long requests."
//       )
//       .catch((err) => {
//         console.error(err);
//       });
//   }
// });

// Chat bot logic end

client.handleEvents();
client.handleCommands();
client.login(token);
