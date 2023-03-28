# Text to Speech

Chloe uses [Google Cloud's Text-to-Speech](https://cloud.google.com/text-to-speech) to generate audio files from text.

_Note: For Brazilian Portuguese, `CHLOE_TTS_LANGUAGE_CODE=pt-BR`
and `CHLOE_TTS_VOICE_NAME=pt-BR-Neural2-C` yields good results._

## Configuration

| Environment Variable           | Default Value   | Description                                                      | Options                                                                  |
|--------------------------------|-----------------|------------------------------------------------------------------|--------------------------------------------------------------------------|
| GOOGLE_APPLICATION_CREDENTIALS |                 | Google Cloud credentials file                                    |                                                                          |
| CHLOE_TTS_LANGUAGE_CODE        | en-US           | Language code for the TTS engine                                 | Refer to the [docs](https://cloud.google.com/text-to-speech/docs/voices) |
| CHLOE_TTS_VOICE_NAME           | en-US-Wavenet-F | Voice name for the TTS engine                                    | Refer to the [docs](https://cloud.google.com/text-to-speech/docs/voices) |
| CHLOE_TTS_AUDIO_ENCODING       | MP3             | Audio format, defaults to MP3, others are available but untested | MP3<br/>LINEAR16<br/>OGG_OPUS<br/>MULAW<br/>ALAW                         |
| CHLOE_TTS_SPEAKING_RATE        | 1.0             | Speaking rate for the TTS engine                                 | Between 0.25 and 4.0                                                     |
| CHLOE_TTS_PITCH                | 0.0             | Pitch for the TTS engine                                         | Between -20.0 and 20.0                                                   |
| CHLOE_TTS_VOLUME_GAIN_DB       | 0.0             | Volume gain for the TTS engine in DB                             | Between -96.0 and 16.0                                                   |

