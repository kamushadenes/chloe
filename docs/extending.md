# Extending Chloe

Chloe can be easily extended to support new actions. You merely need to write the code that performs
the action you want and teach Chloe how to detect the action and call it.

# 1. Adding the action

Add a new package
under [react/actions](https://github.com/kamushadenes/chloe/tree/main/react/actions) with the name
of the action, and implement
the [Action](https://github.com/kamushadenes/chloe/blob/main/structs/action.go) interface.

To facilitate things and reduce duplicate code, there's
generator that writes most of the boiler plate. Be sure to add your action there, along with it's
required parameters, and to set skips for functions you might want to write yourself.

Of everything, only the struct itself, the `Execute` and the `GetNotification` methods are required
to be implemented manually.

You can use [tts](https://github.com/kamushadenes/chloe/tree/main/react/actions/tts/tts.go) as a
reference of a simple action, but basically you need to return an array
of [ResponseObject](https://github.com/kamushadenes/chloe/blob/main/structs/response_object.go#L17)
in the `Execute` method, which you can create using `structs.NewResponseObject(<OBJECT_TYPE>)`.

### Generating code

In case you aren't aware, to generate files in Go we use `go generate`. To do so, go to the
root of the project and
run `go generate ./...` and Go will create several files ending in `_gen.go` with the generated
code.

# 2. Register the action

Register the action in
the [actions map](https://github.com/kamushadenes/chloe/blob/main/react/actions/actions.go) using
all the aliases you want to support.

# 3. Teach Chloe how to detect the action

Edit
the [action_detection](https://github.com/kamushadenes/chloe/blob/main/resources/prompts/chatgpt/action_detection.prompt)
prompt to teach her how to detect them. Simply add it to the `COMMANDS` section by following the
existing examples.

# 4. (Optional) Add few-shot examples

You can optionally add few-shot examples to help Chloe detect the action.
See [Few-shot prompting](https://github.com/kamushadenes/chloe/tree/main/resources/prompts/chatgpt).

# 5. Give back to the community

If you think your action is useful for the community, please
consider [opening a pull request](https://github.com/kamushadenes/chloe/blob/main/CONTRIBUTING.md)
to add it to Chloe.
