# CLI

Chloe can also be used as a command line interface. This is useful for testing purposes.

```bash
./chloe <command>
```

## Support

| **Feature**      | **Supported** |
|------------------|---------------|
| Completion       | Yes           |
| Image Generation | Yes           |
| Text-to-Speech   | Yes           |
| Transcription    | No            |

## Commands

The following commands are available:

- **complete** [args] - Complete the given text using OpenAI's ChatGPT
- **generate** [args] - Generate an image using DALL-E
- **tts** [args]- Converts text to speech
- **forget** - Wipe all context and reset the conversation with the bot

## REPL

If you call complete without any arguments, it will start a REPL.

```bash
./chloe complete
```

[complete.webm](https://user-images.githubusercontent.com/242529/226281153-152b77c3-4d1f-4d22-bb04-41a39cdd740b.webm)