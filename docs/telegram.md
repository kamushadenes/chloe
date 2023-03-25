# Telegram

Chloe can be used as a Telegram bot. To use it, you need to create a bot using
the [BotFather](https://t.me/botfather) and set the `CHLOE_TELEGRAM_TOKEN` environment variable to the
token provided by the BotFather.

1. Start a conversation with your bot
2. Just chat.

Chloe should automatically detect what you want to do and respond accordingly, including
images and voice messages. You can also send voice messages and it will convert them to text and
respond accordingly.

The following commands are available:

- **/listmodes** - Retrieve the list of all available bot modes or personalities
- **/forget** - Wipe all context and reset the conversation with the bot
- **/mode** - Instruct the bot to switch to a different mode or personality
- **/generate** - Generate an image using DALL-E
- **/tts** - Converts text to speech

## Examples

Here, Chloe understands that I want an image of a cat and responds with one. It also understands
when I ask it to give the cat a cute yellow hat, and it preserves the context of the previous image.

![show me a picture of a cat](../.github/resources/images/telegram.png)

You can see how Chloe reasons about the context and understands what I want to do.

![log](../.github/resources/images/log.png)

Here, Chloe says "Beware of the dog, the cat is shady too". I then ask (in a voice message) for it
to show me a picture of what she just said, and she generates relevant images.

![show me a picture of what you just said](../.github/resources/images/telegram2.png)

See
the [Configuration](https://github.com/kamushadenes/chloe/blob/main/docs/configuration.md#telegram)
for more details.
