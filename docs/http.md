# HTTP

Chloe can be used as an HTTP server. By default, it will listen on port `8080`. You can change
this by setting the `PORT` environment variable.

## Authentication

To interact with the HTTP API, you'll need to provide an `Authorization` header with a valid Bearer token.

To create a token, you can run the `create-api-key` command for your user.


First, get your user ID:

```
$ ./chloe list-users

╭───┬────────────┬───────────┬───────────────┬─────────┬───────────────────────────────────╮
│ # │ FIRST NAME │ LAST NAME │ USERNAME      │ MODE    │            EXTERNAL IDS           │
│   │            │           │               │         ├──────────────┬────────────────────┤
│   │            │           │               │         │ INTERFACE    │ ID                 │
├───┼────────────┼───────────┼───────────────┼─────────┼──────────────┼────────────────────┤
│ 1 │ Kamus      │ Hadenes   │ Kamus Hadenes │ default │ discord      │ XPTO               │
├───┼────────────┼───────────┼───────────────┤         ├──────────────┼────────────────────┤
│ 2 │ User       │ CLI       │ cli           │         │ cli          │ cli                │
╰───┴────────────┴───────────┴───────────────┴─────────┴──────────────┴────────────────────╯
```

Take note of the `#` column, which is the user ID.

If your user isn't there, you can talk to the bot through any interface except CLI and it will be created automatically.

Otherwise, you can use the `create-user` command to create a new user.

Then, create a new API key for your user:
```bash
$ ./chloe create-api-key <user_id>

6d011998e67fa56494d822abe4e4b4c22838fa5488c60fdb05488fcb42365606

```

You can now use this token to interact with the HTTP API by providing it in the `Authorization` header like so:

```
Authorization: Bearer <token>
```

## Endpoints

The following endpoints are available:

### POST /api/complete

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

## HTTP Configuration

| Environment Variable | Default Value | Description      | Options |
|----------------------|---------------|------------------|---------|
| CHLOE_HTTP_HOST      | 0.0.0.0       | HTTP server host |         |
| CHLOE_HTTP_PORT      | 8080          | HTTP server port |         |
