# Slack

Chloe can be used as a Slack bot. To use it, you need to create a bot and set
both the `CHLOE_SLACK_TOKEN` and `CHLOE_SLACK_APP_LEVEL_TOKEN` environment variables.

Chloe uses Slack [Socket Mode](https://api.slack.com/apis/connections/socket), which means you don't
need to expose your bot to the internet. The provided manifest will create a bot with the required
scopes and configurations.

1. Create your app at https://api.slack.com/apps using the
   provided [manifest](https://github.com/kamushadenes/chloe/blob/main/docs/files/slack_manifest.json).
2. Copy the bot token and the app level token to the environment variables.
3. Just chat.

Chloe should automatically detect what you want to do and respond accordingly, including generating
images and voice messages.

## Support

| **Feature**      | **Supported** |
|------------------|---------------|
| Completion       | Yes           |
| Image Generation | Yes           |
| Text-to-Speech   | Yes           |
| Transcription    | No            |

## Commands

The following commands are available via DM:

- **/forget** - Wipe all context and reset the conversation with the bot
- **/generate** - Generate an image using DALL-E
- **/tts** - Converts text to speech

## Configuration

| Environment Variable                | Default Value   | Description                                                                                                                                         | Options            |
|-------------------------------------|-----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| CHLOE_SLACK_TOKEN                   |                 | Slack bot token                                                                                                                                     |                    |
| CHLOE_SLACK_APP_LEVEL_TOKEN         |                 | Slack bot app level token                                                                                                                           |                    |
| CHLOE_SLACK_IMAGE_COUNT             | 4               | Number of images to generate when the user asks for an image                                                                                        | Between 1 and 10   |
| CHLOE_SLACK_ONLY_MENTION            | true            | Whether the bot should only respond to mentions                                                                                                     | true<br/>false     |
| CHLOE_SLACK_STREAM_MESSAGES         | false           | Whether to stream messages as they are generated (not recommended)                                                                                  | true<br/>false     |
| CHLOE_SLACK_STREAM_FLUSH_INTERVAL   | 500ms           | Interval between flushing the stream buffer                                                                                                         |                    |
| CHLOE_SLACK_SEND_PROCESSING_MESSAGE | false           | Whether to send a processing message placeholder while the bot is generating it's response, defaults to true if CHLOE_SLACK_STREAM_MESSAGES is true | true<br/>false     |
| CHLOE_SLACK_PROCESSING_MESSAGE      | ↻ Processing... | Message to send as a placeholder while the bot is generating it's response                                                                          |                    |
| CHLOE_SLACK_MAX_MESSAGE_LENGTH      | 2000            | Maximum length of a message                                                                                                                         | Between 1 and 2000 |