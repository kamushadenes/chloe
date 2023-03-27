# Actions

Using [ReAct](https://react-lm.github.io), Chloe can perform actions based on the context of the
conversation.

All actions are defined in the `react` package.

## Supported Actions

### [Google](https://github.com/kamushadenes/chloe/blob/main/react/actions/google/)

Searches Google for information. Results are then scraped and summarized.

### [Image](https://github.com/kamushadenes/chloe/blob/main/react/actions/image/)

Uses OpenAI's [DALL-E](https://openai.com/product/dall-e-2) to generate images, optionally improving
the prompt.

### [Math](https://github.com/kamushadenes/chloe/blob/main/react/actions/math/)

Performs calculations using [govaluate](https://github.com/Knetic/govaluate)

### [News](https://github.com/kamushadenes/chloe/blob/main/react/actions/news/)

Searches news articles using either Google or [NewsAPI](https://newsapi.org) and summarizes them.

### [Scrape](https://github.com/kamushadenes/chloe/blob/main/react/actions/scrape/)

Scrapes a website and adds it to the context.

### [Transcribe](https://github.com/kamushadenes/chloe/blob/main/react/actions/transcribe/)

Transcribes an audio file using Whisper.

### [Wikipedia](https://github.com/kamushadenes/chloe/blob/main/react/wikipedia/)

Searches Wikipedia for information. Results are then scraped and summarized.

### [YouTube Summarizer](https://github.com/kamushadenes/chloe/blob/main/react/youtube_summarizer/)

Summarizes a YouTube video by first downloading it using youtube-dl, then using Whisper to
transcribe it, and finally using a summarization prompt.


