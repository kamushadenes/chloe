# ReAct: YouTube Summarizer

This action allows you to summarize YouTube videos. It works by downloading the video using
youtube-dl (with aria2 optionally), extracting the audio using ffmpeg and then
using [Whisper](https://openai.com/research/whisper) to transcribe the audio.

The transcription is then fed into a summarization prompt to generate a summary of the video, which
is saved for context and returned to the user.

If aria2 is enabled but not available, it will fall back to using youtube-dl without aria2.

## Configuration

| Environment Variable          | Default Value | Description                                                  | Options       |
|-------------------------------|---------------|--------------------------------------------------------------|---------------|
| CHLOE_REACT_YOUTUBE_USE_ARIA2 | true          | Whether to use aria2 for downloading YouTube videos (faster) | true<br>false |