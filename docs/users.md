# User Management

To be able to actually maintain a conversation, Chloe stores user messages and their responses in a
database. Those messages are passed in each request up until the model's request limit, at which
point they are removed from the request from first to last.

Messages are also summarized every few seconds, so we can keep a longer conversation history by
using less tokens.

Those messages are all associated with a user, which is identified by a unique ID. Users are created
automatically upon first interaction and usually don't need to be managed manually.

Still, there are some cases where you may want to manage users manually, such as deleting them (and
their messages), merging users that came from different interfaces (e.g. Telegram and Discord) into
one,
or [adding API keys to users for use with HTTP](https://github.com/kamushadenes/chloe/blob/main/docs/http.md#authentication).

User management is done through the command line.

## Creating users

```
Usage: chloe create-user

Create a new user

Flags:
  -h, --help                 Show context-sensitive help.
      --version              Print version information and quit

  -u, --username=STRING
  -f, --first-name=STRING
  -l, --last-name=STRING
```

## Deleting users

```
Usage: chloe delete-user

Delete a user

Flags:
  -h, --help            Show context-sensitive help.
      --version         Print version information and quit

  -u, --user-id=UINT
```

## Listing users

```
Usage: chloe list-users

List users

Flags:
  -h, --help              Show context-sensitive help.
      --version           Print version information and quit

      --format="table"    Output format, one of: table, markdown
```

## Merging users

Merging users is useful when you want a user to share history between different platforms (e.g.
Telegram and Discord). Chloe will maintain a single conversation history this way.

Use the `list-users` command to find the IDs of the users you want to merge.

```
Usage: chloe merge-users <users> ...

Merge users

Arguments:
  <users> ...    Users to merge

Flags:
  -h, --help       Show context-sensitive help.
      --version    Print version information and quit
```

## Adding external IDs to users

This should hardly ever be necessary, as users are created automatically and their external IDs set
according to the interface they are coming from.

Still, the command is available.

```
Usage: chloe add-external-id

Add external ID to user

Flags:
  -h, --help                  Show context-sensitive help.
      --version               Print version information and quit

  -u, --user-id=UINT
  -e, --external-id=STRING
  -i, --interface=STRING
```

## Deleting external IDs from users

```
Usage: chloe delete-external-id

Delete external ID from user

Flags:
  -h, --help                  Show context-sensitive help.
      --version               Print version information and quit

  -u, --user-id=UINT
  -e, --external-id=STRING
  -i, --interface=STRING
```

## Listing user messages

```
Usage: chloe list-messages <user-id>

List messages

Arguments:
  <user-id>

Flags:
  -h, --help              Show context-sensitive help.
      --version           Print version information and quit

      --format="table"    Output format, one of: table, markdown
```