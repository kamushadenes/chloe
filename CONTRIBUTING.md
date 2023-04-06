# Welcome to Chloe contributing guide <!-- omit in toc -->

Thank you for investing your time in contributing to our project!

Read our [Code of Conduct](./CODE_OF_CONDUCT.md) to keep our community approachable and respectable.

In this guide you will get an overview of the contribution workflow from opening an issue, creating
a PR, reviewing, and merging the PR.

## Getting started

### Issues

#### Create a new issue

If you spot a problem with the
docs, [search if an issue already exists](https://docs.github.com/en/github/searching-for-information-on-github/searching-on-github/searching-issues-and-pull-requests#search-by-the-title-body-or-comments).
If a related issue doesn't exist, you can open a new issue using a
relevant [issue form](https://github.com/kamushadenes/chloe/issues/new/choose).

#### Solve an issue

### Make Changes

#### Make changes locally

1. Fork the repository.

2. Install or update **Golang**

3. Create a working branch and start with your changes!

### Commit your update

Commit the changes once you are happy with them.

#### Linting (optional)

We use [pre-commit](https://pre-commit.com) to lint and format our code.

Installing it is optional, as the CI will run the linters for you, but it's recommended to install
it locally.

You can install it by running:

```bash
pip install pre-commit
```

or

```bash
brew install pre-commit
```

Then run:

```bash
pre-commit install
```

This will install the pre-commit hook in your local repository. Now, every time you commit, the
pre-commit hook will run and check your code for linting and formatting errors.

### Pull Request

When you're finished with the changes, create a pull request, also known as a PR.

- Fill the required fields in the PR template.
- Don't forget to link PR to issue if you are solving one.
- We may ask for changes to be made before a PR can be merged, either
  using [suggested changes](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/incorporating-feedback-in-your-pull-request)
  or pull request comments. You can apply suggested changes directly through the UI. You can make
  any other changes in your fork, then commit them to your branch.
- As you update your PR and apply changes, mark each conversation
  as [resolved](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/commenting-on-a-pull-request#resolving-conversations).
- If you run into any merge issues, checkout
  this [git tutorial](https://github.com/skills/resolve-merge-conflicts) to help you resolve merge
  conflicts and other issues.

### Your PR is merged

Congratulations :tada::tada: The Chloe team thanks you :sparkles:.
