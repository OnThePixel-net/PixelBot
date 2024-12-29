package commands

import (
	"strings"

	"github.com/Paranoia8972/PixelBot/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

func AutoRoleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if len(i.ApplicationCommandData().Options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No subcommand provided.",
			},
		})
		return
	}

	switch i.ApplicationCommandData().Options[0].Name {
	case "add":
		AutoRoleAddCommand(s, i)
	case "get":
		AutoRoleGetCommand(s, i)
	case "remove":
		AutoRoleRemoveCommand(s, i)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Unknown subcommand.",
			},
		})
	}
}

func AutoRoleAddCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options[0].Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please provide a role to add.",
				Flags:   64,
			},
		})
		return
	}

	roleID := options[0].RoleValue(s, "").ID

	err := utils.AddAutoRole(i.GuildID, roleID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to add auto role.",
				Flags:   64,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Auto role added successfully!",
			Flags:   64,
		},
	})
}

func AutoRoleGetCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	autoRoles, err := utils.GetAutoRoles(i.GuildID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to retrieve auto roles.",
				Flags:   64,
			},
		})
		return
	}

	var roleMentions []string
	for _, roleID := range autoRoles {
		roleMentions = append(roleMentions, "<@&"+roleID+">")
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: 64,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Auto Roles",
					Description: "Auto roles for this server.",
					Color:       0x248045,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Roles",
							Value:  "• " + strings.Join(roleMentions, "\n• "),
							Inline: true,
						},
					},
				},
			},
		},
	})
}

func AutoRoleRemoveCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options[0].Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please provide a role to remove.",
				Flags:   64,
			},
		})
		return
	}

	roleID := options[0].RoleValue(s, "").ID

	err := utils.RemoveAutoRole(i.GuildID, roleID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to remove auto role.",
				Flags:   64,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Auto role removed successfully!",
			Flags:   64,
		},
	})
}
