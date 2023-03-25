# CLI

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