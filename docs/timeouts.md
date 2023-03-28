# Timeouts

Chloe has a number of timeouts that can be configured. These are all set to sensible defaults, but
you may want to change them depending on your use case.

| Environment Variable           | Default Value | Description                           | Options |
|--------------------------------|---------------|---------------------------------------|---------|
| CHLOE_TIMEOUT_COMPLETION       | 60s           | Timeout for completion requests       |         |
| CHLOE_TIMEOUT_CHAIN_OF_THOUGHT | 60s           | Timeout for chain of thought analysis |         |
| CHLOE_TIMEOUT_TRANSCRIPTION    | 120s          | Timeout for transcription requests    |         |
| CHLOE_TIMEOUT_MODERATION       | 60s           | Timeout for moderation requests       |         |
| CHLOE_TIMEOUT_IMAGE_GENERATION | 120s          | Timeout for image generation requests |         |
| CHLOE_TIMEOUT_IMAGE_EDIT       | 120s          | Timeout for image edit requests       |         |
| CHLOE_TIMEOUT_IMAGE_VARIATION  | 120s          | Timeout for image variation requests  |         |
| CHLOE_TIMEOUT_TTS              | 10s           | Timeout for TTS requests              |         |
| CHLOE_TIMEOUT_SLOWNESS_WARNING | 5s            | Timeout for slowness warning messages |         |
