> ⚠️ AI is advancing incredibly fast.
> 
> Most of what Chloe and other assistant projects did is now superseeded by advances not only from OpenAI's API but the entire AI space as a whole, including open source models.
> 
> Chloe is no longer needed.
>
> It was incredibly fun to play around with this exciting technology, and while Chloe is no more (for now), I'll keep busy with AI-related projects for a long time.
> 
> Keep hacking!

<hr>

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
with [Google's Text-to-Speech](https://cloud.google.com/text-to-speech) engine and [ElevenLabs' Text-To-Speech](https://elevenlabs.io/) engine to provide versatile
and comprehensive assistance.

It offers multiple interfaces and is able to
understand and respond to complex instructions making use of several tools.

## 🚀 Features

- Calculates and logs the cost of each request as well as the total cost of the session
- Scrapes websites to have them on its context
- Searches Google for information
- Searches and summarizes news articles
- Performs calculations
- Uses Google's Text-to-Speech engine to speak
- Uses OpenAI's DALL-E to generate images
- Automatically summarizes messages in order to have a longer context
- Automatically moderates message using
  OpenAI's [moderation endpoint](https://platform.openai.com/docs/guides/moderation)
- Many more, check [Actions](https://github.com/kamushadenes/chloe/wiki/ReAct-Actions)

[complete.webm](https://user-images.githubusercontent.com/242529/226281153-152b77c3-4d1f-4d22-bb04-41a39cdd740b.webm)

## 🌟 Power Chloe's Growth: Lend Your Support

For the price of a single coffee, you can play a crucial role in advancing Chloe, an exciting
project that's pushing the boundaries of fully autonomous AI! This labor of love is developed during
my precious free time, requiring countless hours of dedication. Your generous contributions keep the
momentum going and are deeply appreciated!

To support Chloe's growth, click [here](https://github.com/sponsors/kamushadenes).

## 📱 Supported Interfaces

- [Command Line (CLI)](https://github.com/kamushadenes/chloe/wiki/Interface-CLI)
- [HTTP REST](https://github.com/kamushadenes/chloe/wiki/Interface-HTTP)
- [Discord](https://github.com/kamushadenes/chloe/wiki/Interface-Discord)
- [Slack](https://github.com/kamushadenes/chloe/wiki/Interface-Slack)
- [Telegram](https://github.com/kamushadenes/chloe/wiki/Interface-Telegram)
- [iOS Shortcut](https://github.com/kamushadenes/chloe/wiki/Interface-iOS) (kinda)

## 🛣️ Roadmap

You can track Chloe's progress on
the [Roadmap project](https://github.com/users/kamushadenes/projects/1).

## 🔗 Dependencies

| **Dependency**                            | **Description**                                                                                                                                                                                                                                                             | **License** | **Type** |
|:------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------------|:---------|
| [aria2](https://aria2.github.io)          | Chloe uses [aria2](https://aria2.github.io/) to speed up the download YouTube videos for transcription. Although highly recommended, this is not a mandatory dependency. If you don't have it installed, Chloe will fall back to using the `youtube-dl` default downloader. | GPL-2.0     | Runtime  |
| [ffmpeg](https://ffmpeg.org)              | Chloe uses [ffmpeg](https://ffmpeg.org/) to convert YouTube videos to audio, and also to convert audio received from Telegram to an appropriate format for Whisper. It is also used to perform cost calculation on Whisper requests.                                        | LGPL-2.1    | Runtime  |
| [imagemagick](https://imagemagick.org)    | Chloe uses [imagemagick](https://imagemagick.org/index.php) to convert images to the appropriate format for DALL-E.                                                                                                                                                         | Apache-2.0  | Runtime  |
| [youtube-dl](https://youtube-dl.org)      | Chloe uses [youtube-dl](https://youtube-dl.org/) to download YouTube videos for transcription.                                                                                                                                                                              | Unlicense   | Runtime  |

## 💾 Installation

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

## ⚙️ Usage

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

## 📚 Documentation

Check out the [Wiki](https://github.com/kamushadenes/chloe/wiki) for more information.

### 🔧 Configuration

See [Configuration](https://github.com/kamushadenes/chloe/wiki/Configuration) for more information.

### 💬 Improving responses

See [Few-shot prompting](https://github.com/kamushadenes/chloe/wiki/Few-shot_Prompting).

### 🛠️ Extending capabilities

See [Extending Chloe](https://github.com/kamushadenes/chloe/wiki/How_to_add_a_new_Action).

### 👥 User management

See [User management](https://github.com/kamushadenes/chloe/wiki/Managing_Users).

## 🤝 Contributing

We welcome contributions! If you would like to improve Chloe, please check out
the [Contributing Guide](CONTRIBUTING.md).

### ⚠️ Disclaimer

There are some rushed parts of the code, and some parts that are not very well documented. You can
also find a few TODOs scattered around the codebase and some duplicate code too.

Feel free to open an issue if you have any questions, or even better, open a pull request!

## 📄 License

Chloe is licensed under the [GPL-3.0 License](LICENSE.md).

## 📚 References

- [ReAct: Synergizing Reasoning and Acting in Language Models](https://react-lm.github.io)
- [A simple Python implementation of the ReAct pattern for LLMs](https://til.simonwillison.net/llms/python-react-pattern)

## 🙏 Acknowledgements

- [pkoukk](https://github.com/pkoukk/tiktoken-go) for the Pure-Go tokenizer
- [sashabaranov](https://github.com/sashabaranov/go-openai) for the Go OpenAI SDK
