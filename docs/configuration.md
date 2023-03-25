# Configuration

There are several defaults defined in the `config` package, including things like database
connection string, generated image sizes and timeouts. Those can be overridden by setting the
relevant environment variables.

Of those, the only required one is the `OPENAI_API_KEY` environment variable. You can get one by
signing up for an account at [OpenAI](https://platform.openai.com/).

## Database

| Environment Variable | Default Value | Description                                                                                           | Options                                     |
|----------------------|---------------|-------------------------------------------------------------------------------------------------------|---------------------------------------------|
| CHLOE_DB_DRIVER      | sqlite        | Database driver to use                                                                                | postgres<br/>mysql<br/>sqlite<br/>sqlserver |
| CHLOE_DB_DSN         | chloe.db      | Database connection string, refer to the [docs](https://gorm.io/docs/connecting_to_the_database.html) |                                             |
| CHLOE_DB_MAX_IDLE    | 2             | Maximum number of idle connections                                                                    |                                             |
| CHLOE_DB_MAX_OPEN    | 10            | Maximum number of open connections                                                                    |                                             |

## HTTP

| Environment Variable | Default Value | Description      | Options |
|----------------------|---------------|------------------|---------|
| CHLOE_HTTP_HOST      | 0.0.0.0       | HTTP server host |         |
| CHLOE_HTTP_PORT      | 8080          | HTTP server port |         |

## Telegram

| Environment Variable       | Default Value | Description                                                  | Options          |
|----------------------------|---------------|--------------------------------------------------------------|------------------|
| CHLOE_TELEGRAM_TOKEN       |               | Telegram bot token                                           |                  |
| CHLOE_TELEGRAM_IMAGE_COUNT | 4             | Number of images to generate when the user asks for an image | Between 1 and 10 |

## Discord

| Environment Variable       | Default Value | Description                                                  | Options          |
|----------------------------|---------------|--------------------------------------------------------------|------------------|
| CHLOE_DISCORD_TOKEN        |               | Discord bot token                                            |                  |
| CHLOE_DISCORD_IMAGE_COUNT  | 4             | Number of images to generate when the user asks for an image | Between 1 and 10 |
| CHLOE_DISCORD_ONLY_MENTION | true          | Whether the bot should only respond to mentions              | true<br/>false   |

## Google Cloud

| Environment Variable           | Default Value   | Description                                                      | Options                                                                  |
|--------------------------------|-----------------|------------------------------------------------------------------|--------------------------------------------------------------------------|
| GOOGLE_APPLICATION_CREDENTIALS |                 | Google Cloud credentials file                                    |                                                                          |
| CHLOE_TTS_LANGUAGE_CODE        | en-US           | Language code for the TTS engine                                 | Refer to the [docs](https://cloud.google.com/text-to-speech/docs/voices) |
| CHLOE_TTS_VOICE_NAME           | en-US-Wavenet-F | Voice name for the TTS engine                                    | Refer to the [docs](https://cloud.google.com/text-to-speech/docs/voices) |
| CHLOE_TTS_AUDIO_ENCODING       | MP3             | Audio format, defaults to MP3, others are available but untested | MP3<br/>LINEAR16<br/>OGG_OPUS<br/>MULAW<br/>ALAW                         |

## OpenAI

| Environment Variable                 | Default Value          | Description                                                                                                                                                                                                                        | Options                                                                                                         |
|--------------------------------------|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| OPENAI_API_KEY                       |                        | OpenAI API key, required                                                                                                                                                                                                           |                                                                                                                 |
| CHLOE_MAX_TOKENS_GPT3Dot5Turbo       | 4096                   | Maximum number of tokens GPT-3.5 Turbo is capable of holding                                                                                                                                                                       |                                                                                                                 |
| CHLOE_MIN_REPLY_TOKENS_GPT3Dot5Turbo | 500                    | Minimum number of tokens GPT-3.5 Turbo should have available to reply with                                                                                                                                                         |                                                                                                                 |
| CHLOE_MODEL_COMPLETION               | gpt-3.5-turbo          | Model to use for completion requests                                                                                                                                                                                               | Refer to the [docs](https://platform.openai.com/docs/api-reference/completions/create#completions/create-model) |
| CHLOE_MODEL_CHAIN_OF_THOUGHT         | gpt-3.5-turbo          | Model to use for chain of thought analysis                                                                                                                                                                                         | Refer to the [docs](https://platform.openai.com/docs/api-reference/completions/create#completions/create-model) |
| CHLOE_MODEL_SUMMARIZATION            | gpt-3.5-turbo          | Model to use for summarization                                                                                                                                                                                                     | Refer to the [docs](https://platform.openai.com/docs/api-reference/completions/create#completions/create-model) |                              
| CHLOE_MODEL_TRANSCRIPTION            | whisper-1              | Model to use for audio transcription requests                                                                                                                                                                                      | Refer to the [docs](https://platform.openai.com/docs/api-reference/audio/create#audio/create-model)             |
| CHLOE_MODEL_MODERATION               | text-moderation-latest | Model to use for message content moderation requests                                                                                                                                                                               | Refer to the [docs](https://platform.openai.com/docs/api-reference/moderations/create#moderations/create-model) |
| CHLOE_IMAGE_GENERATION_SIZE          | 1024x1024              | Size of generated images                                                                                                                                                                                                           | 256x256<br/>512x512<br/>1024x1024                                                                               |
| CHLOE_IMAGE_EDIT_SIZE                | 1024x1024              | Size of generated image edits                                                                                                                                                                                                      | 256x256<br/>512x512<br/>1024x1024                                                                               |
| CHLOE_IMAGE_VARIATION_SIZE           | 1024x1024              | Size of generated image variations                                                                                                                                                                                                 | 256x256<br/>512x512<br/>1024x1024                                                                               |
| CHLOE_MESSAGES_TO_KEEP_FULL_CONTENT  | 4                      | To increase context, the bot summarizes messages in the background using Extreme TLDR. This setting controls how many of the most recent messages it should keep the full content of in order to provide a better user experience. |                                                                                                                 |

## Timeouts

| Environment Variable           | Default Value | Description                           | Options |
|--------------------------------|---------------|---------------------------------------|---------|
| CHLOE_TIMEOUT_COMPLETION       | 60s           | Timeout for completion requests       |         |
| CHLOE_TIMEOUT_CHAIN_OF_THOUGHT | 60s           | Timeout for chain of thought analysis |         |
| CHLOE_TIMEOUT_TRANSCRIPTION    | 60s           | Timeout for transcription requests    |         |
| CHLOE_TIMEOUT_MODERATION       | 60s           | Timeout for moderation requests       |         |
| CHLOE_TIMEOUT_IMAGE_GENERATION | 120s          | Timeout for image generation requests |         |
| CHLOE_TIMEOUT_IMAGE_EDIT       | 120s          | Timeout for image edit requests       |         |
| CHLOE_TIMEOUT_IMAGE_VARIATION  | 120s          | Timeout for image variation requests  |         |
| CHLOE_TIMEOUT_TTS              | 10s           | Timeout for TTS requests              |         |
| CHLOE_TIMEOUT_SLOWNESS_WARNING | 5s            | Timeout for slowness warning messages |         |