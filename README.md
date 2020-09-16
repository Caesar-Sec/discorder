# discordmessages
Go script for sending messages to discord via webhooks. The code was originally taken from <https://ryanwise.me/blog/recon-upgrade-discord/>, which is worth the read. I just slightly modified it for my prefered formatting for discord. 
The original code can be found [here.](https://github.com/leobeosab/hacks/blob/master/go/discordmessage/discordmessage.go)


# To use

export D_NOTIFICATION_WH=[your webhook url]
go run discordmessage.go "Hello, World!" 

## Or get fancy and install
go install discordmessage.go
and now we can use discordmessage from anywhere
