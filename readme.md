# Integration Guide

Following is the sample code of how to integrate the package into your service

```bash
newMsg := discord.MessageTemplate{
	Content: "New Message",
	Embeds: discord.MessageEmbeds{
		{
			Title: "testing msg",
			Color: 16711680,
		},
		{
			Title: "multiple msg",
			Color: 16776960,
		},
	},
}

cl := discord.NewDiscordClient(logrus.New(), &discord.DiscordConfig{
		AuthToken: "<bot_auth_token>",
})

cl.SendDiscordMessage("<channel_id>", discordgo.MessageSend(newMsg))
```

##### Common colors RBG int value

```
yellow - 16776960
green  - 11403055
red    - 16711680
```
