# ReAct: Image

Image uses Dall-E to generate an image from a text prompt.

## Configuration

| Environment Variable              | Default Value | Description                                                                             | Options        |
|-----------------------------------|---------------|-----------------------------------------------------------------------------------------|----------------|
| CHLOE_REACT_IMPROVE_IMAGE_PROMPTS | false         | Whether to improve image prompts, basically doing a second pass on the prompt generator | true<br/>false |

## Other Configuration

Those configurations are interface specific and can be found in their respective documentation, but
are listed here for convenience.

| Environment Variable       | Default Value | Description                                                  | Options          |
|----------------------------|---------------|--------------------------------------------------------------|------------------|
| CHLOE_TELEGRAM_IMAGE_COUNT | 4             | Number of images to generate when the user asks for an image | Between 1 and 10 |
| CHLOE_DISCORD_IMAGE_COUNT  | 4             | Number of images to generate when the user asks for an image | Between 1 and 10 |

## Examples

`Show me a picture of a cat`

![show me a picture of a cat](https://github.com/kamushadenes/chloe/raw/main/.github/resources/images/cat.png)
