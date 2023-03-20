# Chloe AI Assistant

Chloe is a powerful AI Assistant written in Go that leverages OpenAI
technologies ([ChatGPT](https://openai.com/product/gpt-4), [Whisper](https://openai.com/research/whisper),
and [DALL-E](https://openai.com/product/dall-e-2)) along
with [Google's Text-to-Speech](https://cloud.google.com/text-to-speech) engine to provide versatile
and comprehensive
assistance.

It offers multiple interfaces and utilizes
the [Chain of Thought](https://til.simonwillison.net/llms/python-react-pattern) approach to
understand and respond to complex instructions.

## Features

- Use Chain of Thought to determine actions, falling back to standard completion if no action is
  found
- Scrape websites to have them on its context
- Search Google for information
- Perform calculations
- Use Google's Text-to-Speech engine to speak
- Use OpenAI's DALL-E to generate images

Due to the Chain of Thought approach, Chloe can also be extended with additional capabilities by
simply adding new actions.

# Table of Contents

- Installation
- Usage
    - Telegram
    - HTTP
    - Command Line (CLI)
- Contributing
- License

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
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "Say out loud: Hello, Chloe!"}' http://localhost:8080/api/complete
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

[complete.webm](https://user-images.githubusercontent.com/242529/226279941-d9ed067b-927b-40b1-89b2-3589b0f0bf0e.webm)