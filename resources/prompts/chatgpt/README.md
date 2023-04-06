## Few-shot prompting

Chloe is already quite smart, especially using GPT-4, but you may feel that it's not as smart as
you'd like it to be.

In this case, you can
provide [few-shot](https://github.com/openai/openai-python/blob/main/chatml.md#few-shot-prompting)
examples to improve the responses.

To do that, you need a file named `<prompt_name>.examples`
at the [prompt folder](https://github.com/kamushadenes/chloe/tree/main/resources/prompts/chatgpt),
with the same name as
the prompt file (`<prompt_name>.prompt`).

They have the following format:

```
<name> <message>
```

Where name is either `example_user` or `example_assistant`, and `message` is the content of the
message.

```
example_user Hello, how are you?
example_assistant I'm fine, thank you. How are you?
```

By providing a sample conversation and organizing messages in their chronological order, you can
gently steer GPT into the right direction.

Be aware though that this will consume additional tokens at every request, so don't go overboard.
