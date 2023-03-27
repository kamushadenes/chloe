<h1 align="center">Chloe AI Assistant</h1>

<p align="center">A powerful AI assistant</p>

![tests status](https://img.shields.io/github/actions/workflow/status/kamushadenes/chloe/test.yml)
![license](https://img.shields.io/github/license/kamushadenes/chloe)

<img align="left" width=100px src=".github/resources/images/chloe.png" />

[Chloe](https://blog.hadenes.io/post/chloe-ai-assistant/) is a powerful AI Assistant written in Go
that leverages OpenAI
technologies ([ChatGPT](https://openai.com/product/gpt-4),
[Whisper](https://openai.com/research/whisper),
and [DALL-E](https://openai.com/product/dall-e-2)) along
with [Google's Text-to-Speech](https://cloud.google.com/text-to-speech) engine to provide versatile
and comprehensive assistance.

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
- Automatically summarizes messages in order to have a longer context
- Automatically moderates message using
  OpenAI's [moderation endpoint](https://platform.openai.com/docs/guides/moderation)
- Many more, check [Actions](docs/actions.md)

Due to the Chain of Thought approach, Chloe can also be extended with additional capabilities by
simply [adding new actions](https://github.com/kamushadenes/chloe/blob/main/react/react.go#L136).

[complete.webm](https://user-images.githubusercontent.com/242529/226281153-152b77c3-4d1f-4d22-bb04-41a39cdd740b.webm)

## Supported Interfaces

- [Telegram](docs/telegram.md)
- [Discord](docs/discord.md)
- [HTTP REST](docs/http.md)
- [Command Line (CLI)](docs/cli.md)

## Roadmap

- [x] Add additional storage backends
- [x] Add Discord interface
- [ ] Add Slack interface
- [ ] Take action when content is flagged by the moderation
- [ ] Add GPT-4 support (waiting for the API to be released)
- [ ] Support the newly announced [GPT-4 plugins](https://openai.com/blog/chatgpt-plugins)
- [ ] Add more
  actions ([give me ideas!](https://github.com/kamushadenes/chloe/issues/new?assignees=kamushadenes&labels=feature&template=feature_request.md&title=%5BFEATURE%5D+))

I also plan to release an Alexa open hardware clone that will be able to run Chloe, but this will
take quite some time.

# Table of Contents

- [Dependencies](#dependencies)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [References](#references)
- [Acknowledgements](#acknowledgements)

## Dependencies

- aria2
- ffmpeg
- imagemagick
- youtube-dl

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

## Usage

Setup the required environment variables

```bash
export CHLOE_TELEGRAM_TOKEN="your_telegram_bot_token"
export CHLOE_DISCORD_TOKEN="your_discord_bot_token"
export OPENAI_API_KEY="your_openai_api_key"
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/google/application/credentials.json"
```

Running the `chloe` binary will start the bot.

```bash
./chloe
```

## Configuration

See [docs/configuration.md](docs/configuration.md) for more information.

## Contributing

We welcome contributions! If you would like to improve Chloe, please check out
the [Contributing Guide](CONTRIBUTING.md).

## License

Chloe is licensed under the [MIT License](LICENSE.md).

## References

- [ReAct: Synergizing Reasoning and Acting in Language Models](https://react-lm.github.io)
- [A simple Python implementation of the ReAct pattern for LLMs](https://til.simonwillison.net/llms/python-react-pattern)

## Acknowledgements

- [sashabaranov](https://github.com/sashabaranov/go-openai) for the Go OpenAI SDK
- [awesome-chatgpt-prompts](https://github.com/f/awesome-chatgpt-prompts) for the personalities
