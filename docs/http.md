# HTTP

Chloe can also be used as an HTTP server. By default, it will listen on port `8080`. You can change
this by setting the `PORT` environment variable.

See
the [Configuration](https://github.com/kamushadenes/chloe/blob/main/docs/configuration.md#http)
for more details.

The following endpoints are available:

#### POST /api/complete

This endpoint will complete the given text using OpenAI's ChatGPT.

**Request**

```json
{
  "content": "Hello, Chloe!"
}
```

**Response**

#### Text

If the response is a text (the default), you'll receive a streamed response. You can use curl's `-N`
flag to receive the text as it's generated.

```bash
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "Hello, Chloe!"}' http://localhost:8080/api/complete
```

```
Hello Henrique! How can I assist you today?
```

#### Image

In case the Chain of Thought detects the user wants to generate an image, you'll receive the PNG
image as a response.

```bash
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "show me a picture of a cat"}' http://localhost:8080/api/complete
```

![show me a picture of a cat](../.github/resources/images/cat.png)

#### Audio

In case the Chain of Thought detects the user wants to convert text to speech, you'll receive the
MP3 audio as a response.

```bash
curl -N -X POST -H "Content-Type: application/json" -d '{"content": "Say out loud: Hello, my name is Chloe!"}' http://localhost:8080/api/complete
```

https://user-images.githubusercontent.com/242529/226274960-9086191d-4267-476e-8c0f-f3a449bfac53.mp4