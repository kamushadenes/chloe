# Discord

Chloe can be used as a Discord bot. To use it, you need to create a bot and set
the `CHLOE_DISCORD_TOKEN` environment variable.

1. Invite your bot to a server (try [this page](https://discordapi.com/permissions.html))
2. Just chat.

Chloe should automatically detect what you want to do and respond accordingly, including
images and voice messages. You can also send voice messages and it will convert them to text and
respond accordingly.

The following commands are available via DM:

- **/forget** - Wipe all context and reset the conversation with the bot
- **/generate** - Generate an image using DALL-E
- **/tts** - Converts text to speech

See
the [Configuration](https://github.com/kamushadenes/chloe/blob/main/docs/configuration.md#discord)
for more details.