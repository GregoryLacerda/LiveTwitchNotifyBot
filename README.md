A simple bot in golang that notify on selected discor channel id, if streamers has change the game they are playing in twitch.

To use this bot you will need to create a discord bot, add to your server and channel and get Bot token and channel id
Create a developer in twitch (https://dev.twitch.tv/) and has in hands client id and client secret.

with all that, just add a .env file and run project. see how this .env file need to be:

TWITCH_CLIENT_ID=<CLIENT-ID>
TWITCH_CLIENT_SECRET=<YUOR-CLIENT-SECRET>
DISCORD_BOT_TOKEN=<YOUR-DISCORD-BOT>
DISCORD_BOT_CHANNEL_ID=<YOUR-CHANNEL-ID>
TWITCH_CHANNELS_TO_MONITOR=user_login=yoda&user_login=gaules
