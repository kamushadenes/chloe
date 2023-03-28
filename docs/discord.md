# Discord

Chloe can be used as a Discord bot. To use it, you need to create a bot and set
the `CHLOE_DISCORD_TOKEN` environment variable.

1. Invite your bot to a server (try [this page](https://discordapi.com/permissions.html))
2. Just chat.

Chloe should automatically detect what you want to do and respond accordingly, including
images and voice messages. You can also send voice messages and it will convert them to text and
respond accordingly.

## Commands

The following commands are available via DM:

- **/forget** - Wipe all context and reset the conversation with the bot
- **/generate** - Generate an image using DALL-E
- **/tts** - Converts text to speech

## Configuration

| Environment Variable                        | Default Value   | Description                                                                                                                                           | Options            |
|---------------------------------------------|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| CHLOE_DISCORD_TOKEN                         |                 | Discord bot token                                                                                                                                     |                    |
| CHLOE_DISCORD_IMAGE_COUNT                   | 4               | Number of images to generate when the user asks for an image                                                                                          | Between 1 and 10   |
| CHLOE_DISCORD_ONLY_MENTION                  | true            | Whether the bot should only respond to mentions                                                                                                       | true<br/>false     |
| CHLOE_DISCORD_RANDOM_STATUS_UPDATE_INTERVAL | 1m              | Interval between random status updates, set to 0 to disable                                                                                           |                    |
| CHLOE_DISCORD_STREAM_MESSAGES               | false           | Whether to stream messages as they are generated (not recommended)                                                                                    | true<br/>false     |
| CHLOE_DISCORD_STREAM_FLUSH_INTERVAL         | 500ms           | Interval between flushing the stream buffer                                                                                                           |                    |
| CHLOE_DISCORD_SEND_PROCESSING_MESSAGE       | false           | Whether to send a processing message placeholder while the bot is generating it's response, defaults to true if CHLOE_DISCORD_STREAM_MESSAGES is true | true<br/>false     |
| CHLOE_DISCORD_PROCESSING_MESSAGE            | ↻ Processing... | Message to send as a placeholder while the bot is generating it's response                                                                            |                    |
| CHLOE_DISCORD_MAX_MESSAGE_LENGTH            | 2000            | Maximum length of a message                                                                                                                           | Between 1 and 2000 |