package commands

import (
	"time"

	"github.com/Paranoia8972/PixelBot/internal/app/config"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

var cfg *config.Config

func init() {
	cfg = config.LoadConfig()
}

func RespondWithMessage(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   64,
		},
	})
}

func RegisterCommands(s *discordgo.Session, cfg *config.Config) {

	Commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Responds with the Bot's latency.",
		},
		{
			Name:        "clear",
			Description: "Deletes messages from a channel.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "count",
					Description: "Number of messages to delete.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
				},
				{
					Name:        "user",
					Description: "User whose messages to delete.",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    false,
				},
			},
		},
		{
			Name:        "radio",
			Description: "Controls the radio.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "start",
					Description: "Starts the radio.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "stop",
					Description: "Stops the radio.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:                     "role",
			Description:              "Role management",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageRoles}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "add",
					Description: "Add a role to a member",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "user",
							Description: "User to add the role to",
							Type:        discordgo.ApplicationCommandOptionUser,
							Required:    true,
						},
						{
							Name:        "role",
							Description: "Role to add",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
				{
					Name:        "remove",
					Description: "Remove a role from a member",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "user",
							Description: "User to remove the role from",
							Type:        discordgo.ApplicationCommandOptionUser,
							Required:    true,
						},
						{
							Name:        "role",
							Description: "Role to remove",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
				{
					Name:        "addall",
					Description: "Add a role to all members",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "role",
							Description: "Role to add to all members",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
				{
					Name:        "removeall",
					Description: "Remove a role from all members",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "role",
							Description: "Role to remove from all members",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:                     "say",
			Description:              "Repeats a message.",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageMessages}[0],
		},
		{
			Name:                     "edit",
			Description:              "Edits a message.",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageMessages}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "message_id",
					Description: "ID of the message to edit.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:                     "welcome",
			Description:              "Manage welcome channel",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionAdministrator}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "set",
					Description: "Set the welcome channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "channel",
							Description: "Channel to set as welcome channel",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
						{
							Name:        "message",
							Description: "Welcome message",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
				{
					Name:        "get",
					Description: "Get the current welcome channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "delete",
					Description: "Delete the welcome channel entry",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:                     "social",
			Description:              "Manage social updates channel and accounts",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionAdministrator}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "set",
					Description: "Set the social updates channel and accounts",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "channel",
							Description: "Channel to set as social updates channel",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
						{
							Name:        "youtube",
							Description: "YouTube username",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    false,
						},
						{
							Name:        "twitch",
							Description: "Twitch username",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    false,
						},
						{
							Name:        "twitter",
							Description: "Twitter username",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    false,
						},
					},
				},
				{
					Name:        "get",
					Description: "Get the current social updates channel and accounts",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "delete",
					Description: "Delete the social updates channel entry",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:                     "ticket",
			Description:              "Manage ticket system",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionAdministrator}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "setup",
					Description: "Setup the ticket system",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name: "channel",

							Description: "Channel to set as ticket channel",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
						{
							Name:        "category",
							Description: "Category to set as ticket category",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
						{
							Name:        "transcript",
							Description: "Channel to send transcripts to",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
					},
				},
				{
					Name:        "send",
					Description: "Sends a new message with a button to create a ticket",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:                     "giveaway",
			Description:              "Manage giveaways",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageMessages}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "start",
					Description: "Start a new giveaway",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "duration",
							Description: "Duration of the giveaway (e.g., 1h, 30m)",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
						{
							Name:        "winners",
							Description: "Number of winners",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    true,
						},
						{
							Name:        "prize",
							Description: "Prize of the giveaway",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
				{
					Name:        "end",
					Description: "End an ongoing giveaway",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "message_id",
							Description: "Message ID of the giveaway message",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
				{
					Name:        "reroll",
					Description: "Reroll winners for a giveaway",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "message_id",
							Description: "Message ID of the giveaway message",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:                     "autorole",
			Description:              "Manage auto roles",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageRoles}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "add",
					Description: "Add an auto role",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "role",
							Description: "Role to add",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
				{
					Name:        "get",
					Description: "Get auto roles",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "remove",
					Description: "Remove an auto role",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "role",
							Description: "Role to remove",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:        "mcstatus",
			Description: "Get the status of the Minecraft server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "server_ip",
					Description: "Server to get status for",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:                     "ban",
			Description:              "Ban a user from the server",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionBanMembers}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "The user to ban",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
				{
					Name:        "reason",
					Description: "Reason for the ban",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
				},
			},
		},
		{
			Name:                     "kick",
			Description:              "Kick a user from the server",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionKickMembers}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "The user to kick",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
				{
					Name:        "reason",
					Description: "Reason for the kick",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
				},
			},
		},
		{
			Name:                     "unban",
			Description:              "Unban a user from the server",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionBanMembers}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "user",
					Description:  "User to unban",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		{
			Name:        "coinflip",
			Description: "Flip a coin",
		},
		{
			Name:        "randomnumber",
			Description: "Generate a random number",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "max",
					Description: "Maximum number to generate",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
				},
			},
		},
		{
			Name:        "chooser",
			Description: "Choose a random item from a list or message reactions",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "list",
					Description: "Choose from a comma-separated list",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "items",
							Description: "List of items to choose from, separated by commas",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
						{
							Name:        "amount",
							Description: "Number of items to choose",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    false,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "message",
					Description: "Choose from users who reacted to a message",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "message_id",
							Description: "ID of the message to choose from",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
						{
							Name:        "amount",
							Description: "Number of users to choose",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    false,
						},
					},
				},
			},
		},
		{
			Name:        "version",
			Description: "Get the bot's version",
		},
		{
			Name:        "advent",
			Description: "Advent calendar",
		},
		{
			Name:        "level",
			Description: "View your level or the level of another user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "User to view",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    false,
				},
			},
		},
		{
			Name:                     "leveling",
			Description:              "Manage leveling system",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageChannels}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "set_reward",
					Description: "Set the reward for a level",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "level",
							Description: "Level to set the reward for",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    true,
						},
						{
							Name:        "role",
							Description: "Role to give as a reward",
							Type:        discordgo.ApplicationCommandOptionRole,
							Required:    true,
						},
					},
				},
				{
					Name:        "get_reward",
					Description: "Get the reward for a level",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "remove_reward",
					Description: "Remove the reward for a level",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "level",
							Description: "Level to remove the reward for",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    true,
						},
					},
				},
				{
					Name:        "set_channel_requirement",
					Description: "Set the level requirement for a channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "channel",
							Description: "Channel to set the requirement for",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
						{
							Name:        "requirement",
							Description: "The level required to send messages in the channel",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    true,
						},
					},
				},
				{
					Name:        "get_channel_requirement",
					Description: "Get the level requirement for your server",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "delete_channel_requirement",
					Description: "Delete the level requirement for a channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "channel",
							Description: "Channel to delete the requirement for",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
					},
				},
				{
					Name:        "setlevelchannel",
					Description: "Set the channel for level-up messages",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "channel",
							Description: "Channel for level-up messages",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:                     "status",
			Description:              "Manage bot status messages",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageServer}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "add",
					Description: "Add a new status",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "message",
							Description: "Status message to display",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
						{
							Name:        "type",
							Description: "Activity type",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Required:    true,
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								{Name: "Playing", Value: 0},
								{Name: "Streaming", Value: 1},
								{Name: "Listening", Value: 2},
								{Name: "Watching", Value: 3},
								{Name: "Custom", Value: 4},
								{Name: "Competing", Value: 5},
							},
						},
					},
				},
				{
					Name:        "remove",
					Description: "Remove a status",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "message",
							Description: "Status message to remove",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
				{
					Name:        "list",
					Description: "List all statuses",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:                     "counting",
			Description:              "Manage the counting game channel",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageChannels}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "set",
					Description: "Set the counting channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "channel",
							Description: "Channel to use for counting",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Required:    true,
						},
					},
				},
				{
					Name:        "get",
					Description: "Get the current counting channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "delete",
					Description: "Delete the counting channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:                     "dm",
			Description:              "Send a direct message to a user",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionAdministrator}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "User to send the message to",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
		},
		{
			Name:                     "dmlog",
			Description:              "Set the channel for logging DMs",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionAdministrator}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "channel",
					Description: "Channel to log DMs in",
					Type:        discordgo.ApplicationCommandOptionChannel,
					Required:    true,
				},
			},
		},
		{
			Name:                     "sticky",
			Description:              "Create a sticky message in the current channel",
			DefaultMemberPermissions: &[]int64{discordgo.PermissionManageMessages}[0],
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "create",
					Description: "Create a sticky message",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "list",
					Description: "List all sticky messages",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "remove",
					Description: "Remove the sticky message",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "message_id",
							Description: "ID of the message to make sticky",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	commands := make([]*discordgo.ApplicationCommand, len(Commands))
	copy(commands, Commands)

	start := time.Now()
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, cfg.GuildID, commands)
	duration := time.Since(start)

	if err != nil {
		if rErr, ok := err.(*discordgo.RESTError); ok && rErr.Response.StatusCode == 429 {
			color.Yellow("Rate limited: %v", rErr.Message)
		} else {
			color.Red("Cannot bulk overwrite commands: %v", err)
		}
	} else {
		for _, cmd := range commands {
			color.Blue("Registered command: %s", cmd.Name)
		}
		color.Blue("Registered %d commands in %v", len(commands), duration)
	}
}
