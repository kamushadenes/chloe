<h1 align="center">Chloe</h1>

<p align="center">A powerful AI assistant</p>

<p align="center">
  <img src=".github/resources/images/chloe_avatar.png" alt="Chloe" />
</p>


![tests status](https://img.shields.io/github/actions/workflow/status/kamushadenes/chloe/test.yml?label=tests)
![license](https://img.shields.io/github/license/kamushadenes/chloe)

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

- Uses Chain of Thought to determine actions, falling back to standard completion if no action is
  found
- Scrapes websites to have them on its context
- Searches Google for information
- Searches and summarizes news articles
- Performs calculations
- Uses Google's Text-to-Speech engine to speak
- Uses OpenAI's DALL-E to generate images
- Uses the official OpenAI tokenizer (via Rust bindings) to properly handle token counts
- Automatically summarizes messages in order to have a longer context
- Automatically moderates message using
  OpenAI's [moderation endpoint](https://platform.openai.com/docs/guides/moderation)
- Many more, check [Actions](docs/actions.md)

Due to the Chain of Thought approach, Chloe can also be extended with additional capabilities by
simply [adding new actions](https://github.com/kamushadenes/chloe/blob/main/react/react.go#L136).

[complete.webm](https://user-images.githubusercontent.com/242529/226281153-152b77c3-4d1f-4d22-bb04-41a39cdd740b.webm)

## Supported Interfaces

- [Command Line (CLI)](docs/cli.md)
- [HTTP REST](docs/http.md)
- [Discord](docs/discord.md)
- [Telegram](docs/telegram.md)
- [Slack](docs/slack.md)
- [iOS Shortcut](docs/ios.md) (kinda)

## Roadmap

- [x] Add additional storage backends
- [x] Add Discord interface
- [x] Add Slack interface
- [ ] Take action when content is flagged by the moderation
- [x] Add GPT-4 support
- [ ] Support the newly announced [GPT-4 plugins](https://openai.com/blog/chatgpt-plugins)
- [ ] Add Wolfram Alpha integration
- [x] Add LaTeX rendering support
- [ ] Add more
  actions ([give me ideas!](https://github.com/kamushadenes/chloe/issues/new?assignees=kamushadenes&labels=feature&template=feature_request.md&title=%5BFEATURE%5D+))

I also plan to release an Alexa open hardware clone that will be able to run Chloe, but this will
take quite some time.

## Dependencies

| **Dependency**                            | **Description**                                                                                                                                                                                                                                                             | **License** | **Type** |
|:------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------------|:---------|
| [aria2](https://aria2.github.io)          | Chloe uses [aria2](https://aria2.github.io/) to speed up the download YouTube videos for transcription. Although highly recommended, this is not a mandatory dependency. If you don't have it installed, Chloe will fall back to using the `youtube-dl` default downloader. | GPL-2.0     | Runtime  |
| [cargo](https://doc.rust-lang.org/cargo/) | Chloe uses [cargo](https://doc.rust-lang.org/cargo/) to build the tokenizer bindings. This is only necessary during the build process, so if you're using the pre-built binaries you can skip this dependency.                                                              | Apache-2.0  | Build    |
| [ffmpeg](https://ffmpeg.org)              | Chloe uses [ffmpeg](https://ffmpeg.org/) to convert YouTube videos to audio, and also to convert audio received from Telegram to an appropriate format for Whisper.                                                                                                         | LGPL-2.1    | Runtime  |
| [imagemagick](https://imagemagick.org)    | Chloe uses [imagemagick](https://imagemagick.org/index.php) to convert images to the appropriate format for DALL-E.                                                                                                                                                         | Apache-2.0  | Runtime  |
| [youtube-dl](https://youtube-dl.org)      | Chloe uses [youtube-dl](https://youtube-dl.org/) to download YouTube videos for transcription.                                                                                                                                                                              | Unlicense   | Runtime  |

## Installation

### Pre-built binaries

Pre-built binaries are available for Linux (amd64) and macOS (Intel and M1).

You can download them from the [releases page](https://github.com/kamushadenes/chloe/releases).

Windows might work if you compile from source, but it's untested.

### From source

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
make
```

## Usage

Setup the required environment variables

```bash
# This is the only mandatory variable
export OPENAI_API_KEY="your_openai_api_key"                       

# Only necessary if you want to use the Telegram interface
export CHLOE_TELEGRAM_TOKEN="your_telegram_bot_token"             

# Only necessary if you want to use the Discord interface
export CHLOE_DISCORD_TOKEN="your_discord_bot_token"               

# Only necessary if you want to use the Text-to-Speech engine
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/credentials.json" 
```

Running the `chloe` binary will start the bot.

```bash
cd cmd/chloe/

./chloe
```

Here's the complete list of available commands:

```
Usage: chloe <command>

Chloe is a powerful AI Assistant

Running Chloe without arguments will start the server

Flags:
  -h, --help       Show context-sensitive help.
      --version    Print version information and quit

Commands:
  complete              Complete a prompt
  generate              Generate an prompt
  tts                   Generate an audio from a prompt
  action                Performs an action
  forget                Forget all users
  count-tokens          Count tokens
  create-user           Create a new user
  delete-user           Delete a user
  list-users            List users
  merge-users           Merge users
  add-external-id       Add external ID to user
  delete-external-id    Delete external ID from user
  list-messages         List messages
  create-api-key        Create an API key for use with the HTTP interface

Run "chloe <command> --help" for more information on a command.
```

## Configuration

See [docs/configuration.md](docs/configuration.md) for more information.

## Improving responses

See [Few-shot prompting](https://github.com/kamushadenes/chloe/tree/main/resources/prompts/chatgpt)

## Contributing

We welcome contributions! If you would like to improve Chloe, please check out
the [Contributing Guide](CONTRIBUTING.md).

### Disclaimer

There are some rushed parts of the code, and some parts that are not very well documented. You can
also find a lot of TODOs scattered around the codebase and some duplicated code too.

Feel free to open an issue if you have any questions, or even better, open a pull request!

## License

Chloe is licensed under the [MIT License](LICENSE.md).

## References

- [ReAct: Synergizing Reasoning and Acting in Language Models](https://react-lm.github.io)
- [A simple Python implementation of the ReAct pattern for LLMs](https://til.simonwillison.net/llms/python-react-pattern)

## Acknowledgements

- [Torantulino](https://github.com/Torantulino/Auto-GPT) for some prompt improving ideas
- [j178](https://github.com/j178/tiktoken-go) for the tokenizer bindings
- [sashabaranov](https://github.com/sashabaranov/go-openai) for the Go OpenAI SDK