package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type DiscordConfig struct {
	AuthToken string
}

type discord struct {
	logger *logrus.Logger
	config *DiscordConfig
}

type MessageTemplate discordgo.MessageSend
type MessageEmbeds []*discordgo.MessageEmbed
type MessageFile []*discordgo.File
type MessageEmbedImg *discordgo.MessageEmbedImage

type discordClient interface {
	SendDiscordMessage(channelId string, msgT discordgo.MessageSend)
}

func NewDiscordClient(logger *logrus.Logger, config *DiscordConfig) discordClient {
	return &discord{
		logger: logger,
		config: config,
	}
}

func (d *discord) SendDiscordMessage(channelId string, msg discordgo.MessageSend) {
	goBot, err := discordgo.New("Bot " + d.config.AuthToken)
	if err != nil {
		d.logger.Error("error in initialising ", err)
		return
	}

	_, err = goBot.ChannelMessageSendComplex(channelId, &msg)
	if err != nil {
		d.logger.Error("unable to send message ", err)
		return
	}

	d.logger.Info("Message sent successfully")
}
