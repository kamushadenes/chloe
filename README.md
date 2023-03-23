# Chloe AI Assistant

Chloe is a powerful AI Assistant written in Go that leverages OpenAI
technologies ([ChatGPT](https://openai.com/product/gpt-4), [Whisper](https://openai.com/research/whisper),
and [DALL-E](https://openai.com/product/dall-e-2)) along
with [Google's Text-to-Speech](https://cloud.google.com/text-to-speech) engine to provide versatile
and comprehensive assistance.

It offers multiple interfaces and utilizes
the [Chain of Thought](https://til.simonwillison.net/llms/python-react-pattern) approach to
understand and respond to complex instructions.

[Blog Post](https://blog.hadenes.io/post/chloe-ai-assistant/)

## Features

- Use Chain of Thought to determine actions, falling back to standard completion if no action is
  found
- Scrape websites to have them on its context
- Search Google for information
- Perform calculations
- Use Google's Text-to-Speech engine to speak
- Use OpenAI's DALL-E to generate images
- Automatically summarizes messages in order to have a longer context

Due to the Chain of Thought approach, Chloe can also be extended with additional capabilities by
simply [adding new actions](https://github.com/kamushadenes/chloe/blob/main/react/react.go#L116).

## Roadmap

- [x] Add additional storage backends
- [ ] Add Discord interface
- [ ] Add Slack interface
- [ ] Add more
  actions ([give me ideas!](https://github.com/kamushadenes/chloe/issues/new?assignees=kamushadenes&labels=feature&template=feature_request.md&title=%5BFEATURE%5D+))

I also plan to release an Alexa open hardware clone that will be able to run Chloe, but this will
take quite some time.

# Table of Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Telegram](#telegram)
    - [HTTP](#http)
    - [Command Line (CLI)](#command-line-cli)
- [Contributing](#contributing)
- [License](#license)
- [References](#references)
- [Acknowledgements](#acknowledgements)

## Installation

1. Clone the repository

```bash
git clone https://github.com/kamushadenes/chloe.git
```

2. Change directory to the project folder

```bash
cd chloe
```

3. Compile

```bash
go build -o chloe cmd/chloe/main.go
```

4. Setup the required environment variables

```bash
export TELEGRAM_TOKEN="your_telegram_bot_token"
export OPENAI_API_KEY="your_openai_api_key"
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/google/application/credentials.json"
```

## Usage

Running the `chloe` binary will start the bot.

```bash
./chloe
```

You can use Chloe in multiple ways.

### Telegram

Chloe can be used as a Telegram bot. To use it, you need to create a bot using
the [BotFather](https://t.me/botfather) and set the `TELEGRAM_TOKEN` environment variable to the
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

#### Examples

Here, Chloe understands that I want an image of a cat and responds with one. It also understands
when I ask it to give the cat a cute yellow hat, and it preserves the context of the previous image.

![show me a picture of a cat](.github/resources/images/telegram.png)

You can see how Chloe reasons about the context and understands what I want to do.

![log](.github/resources/images/log.png)

Here, Chloe says "Beware of the dog, the cat is shady too". I then ask (in a voice message) for it
to show me a picture of what she just said, and she generates relevant images.

![show me a picture of what you just said](.github/resources/images/telegram2.png)

### HTTP

Chloe can also be used as an HTTP server. By default, it will listen on port `8080`. You can change
this by setting the `PORT` environment variable.

The following endpoints are available:

#### POST /api/complete

This endpoint will complete the given text using OpenAI's ChatGPT.

**Request**

```json
{
  "content": "Hello, Chloe!"
}
```

**Response**

#### Text

If the response is a text (the default), you'll receive a streamed response. You can use curl's `-N`
flag to receive the text as it's generated.

```bash
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "Hello, Chloe!"}' http://localhost:8080/api/complete
```

```
Hello Henrique! How can I assist you today?
```

#### Image

In case the Chain of Thought detects the user wants to generate an image, you'll receive the PNG
image as a response.

```bash
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "show me a picture of a cat"}' http://localhost:8080/api/complete
```

![show me a picture of a cat](.github/resources/images/cat.png)

#### Audio

In case the Chain of Thought detects the user wants to convert text to speech, you'll receive the
MP3 audio as a response.

```bash
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "Say out loud: Hello, my name is Chloe!"}' http://localhost:8080/api/complete
```

https://user-images.githubusercontent.com/242529/226274960-9086191d-4267-476e-8c0f-f3a449bfac53.mp4

### Command Line (CLI)

Chloe can also be used as a command line interface. This is useful for testing purposes.

```bash
./chloe <command>
```

The following commands are available:

- **complete** [args] - Complete the given text using OpenAI's ChatGPT
- **generate** [args] - Generate an image using DALL-E
- **tts** [args]- Converts text to speech
- **forget** - Wipe all context and reset the conversation with the bot

If you call complete without any arguments, it will start a REPL.

```bash
./chloe complete
```

[complete.webm](https://user-images.githubusercontent.com/242529/226281153-152b77c3-4d1f-4d22-bb04-41a39cdd740b.webm)

## Configuration

There are several defaults defined in the `config` package, including things like database
connection string, generated image sizes and timeouts. Those can be overridden by setting the
relevant environment variables.

Of those, the only required one is the `OPENAI_API_KEY` environment variable. You can get one by
signing up for an account at [OpenAI](https://platform.openai.com/).

### Database

| Environment Variable | Default Value | Description                                                                                           | Options                                     |
|----------------------|---------------|-------------------------------------------------------------------------------------------------------|---------------------------------------------|
| CHLOE_DB_DRIVER      | sqlite        | Database driver to use                                                                                | postgres<br/>mysql<br/>sqlite<br/>sqlserver |
| CHLOE_DB_DSN         | chloe.db      | Database connection string, refer to the [docs](https://gorm.io/docs/connecting_to_the_database.html) |                                             |
| CHLOE_DB_MAX_IDLE    | 2             | Maximum number of idle connections                                                                    |                                             |
| CHLOE_DB_MAX_OPEN    | 10            | Maximum number of open connections                                                                    |                                             |

### HTTP

| Environment Variable | Default Value | Description      | Options |
|----------------------|---------------|------------------|---------|
| CHLOE_HTTP_HOST      | 0.0.0.0       | HTTP server host |         |
| CHLOE_HTTP_PORT      | 8080          | HTTP server port |         |

### Telegram

| Environment Variable       | Default Value | Description                                                  | Options          |
|----------------------------|---------------|--------------------------------------------------------------|------------------|
| CHLOE_TELEGRAM_TOKEN       |               | Telegram bot token                                           |                  |
| CHLOE_TELEGRAM_IMAGE_COUNT | 4             | Number of images to generate when the user asks for an image | Between 1 and 10 |

### Google Cloud

| Environment Variable           | Default Value   | Description                                                      | Options                                                                  |
|--------------------------------|-----------------|------------------------------------------------------------------|--------------------------------------------------------------------------|
| GOOGLE_APPLICATION_CREDENTIALS |                 | Google Cloud credentials file                                    |                                                                          |
| CHLOE_TTS_LANGUAGE_CODE        | en-US           | Language code for the TTS engine                                 | Refer to the [docs](https://cloud.google.com/text-to-speech/docs/voices) |
| CHLOE_TTS_VOICE_NAME           | en-US-Wavenet-F | Voice name for the TTS engine                                    | Refer to the [docs](https://cloud.google.com/text-to-speech/docs/voices) |
| CHLOE_TTS_AUDIO_ENCODING       | MP3             | Audio format, defaults to MP3, others are available but untested | MP3<br/>LINEAR16<br/>OGG_OPUS<br/>MULAW<br/>ALAW                         |

### OpenAI

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

### Timeouts

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

## Contributing

We welcome contributions! If you would like to improve Chloe, please follow these steps:

1. Fork the repository
2. Make your changes
3. Open a pull request

## License

Chloe is licensed under the [MIT License](LICENSE).

## References

- [ReAct: Synergizing Reasoning and Acting in Language Models](https://react-lm.github.io)
- [A simple Python implementation of the ReAct pattern for LLMs](https://til.simonwillison.net/llms/python-react-pattern)

## Acknowledgements

- [sashabaranov](https://github.com/sashabaranov/go-openai) for the Go OpenAI SDK
- [awesome-chatgpt-prompts](https://github.com/f/awesome-chatgpt-prompts) for the personalities
