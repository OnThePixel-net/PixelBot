package commands

import (
	"log"

	"github.com/Paranoia8972/PixelBot/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

func WelcomeCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please provide a subcommand.",
				Flags:   64,
			},
		})
		return
	}

	subCommand := options[0].Name
	switch subCommand {
	case "set":
		channelID := options[0].Options[0].ChannelValue(nil).ID
		welcomeMessage := options[0].Options[1].StringValue()
		guildID := i.GuildID

		err := utils.SetWelcomeChannel(guildID, channelID, welcomeMessage)
		if err != nil {
			log.Printf("Failed to set welcome channel: %v", err)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to set welcome channel.",
					Flags:   64,
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Welcome channel and message set successfully!",
				Flags:   64,
			},
		})
	case "get":
		guildID := i.GuildID
		welcomeChannel, err := utils.GetWelcomeChannel(guildID)
		if err != nil {
			log.Printf("Failed to get welcome channel: %v", err)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to get welcome channel.",
					Flags:   64,
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: 64,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Welcome Channel",
						Description: "Current welcome channel and message",
						Color:       0x248045,
						Image: &discordgo.MessageEmbedImage{
							URL: "https://i.imgur.com/RAClg4Q.png",
						},
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "Channel",
								Value:  "<#" + welcomeChannel.ChannelID + ">",
								Inline: true,
							},
							{
								Name:   "Message",
								Value:  welcomeChannel.WelcomeMessage,
								Inline: true,
							},
						},
					},
				},
			},
		})
	case "delete":
		guildID := i.GuildID
		err := utils.DeleteWelcomeChannel(guildID)
		if err != nil {
			log.Printf("Failed to delete welcome channel: %v", err)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to delete welcome channel.",
					Flags:   64,
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Welcome channel deleted successfully!",
				Flags:   64,
			},
		})
	}
}
