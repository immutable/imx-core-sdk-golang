# Contributing to Immutable X's Golang SDK

First of all, thanks for being here! We greatly appreciate all community feedback and contributions to this repository.

## Alternative ways to get in touch
* Ask usage questions in our [forum](https://forum.immutable.com/).
* Get involved with discussions on the #dev-discussion channel on [Discord](https://discord.com/invite/immutablex).

## Contribution guidelines
### Coding conventions
The code should:
* Have meaningful variable names.
* Have no lint issues.
* Have meaningful slugs, file names, directory names and directory structure.
* Be easy to fix, refactor and scale.

Do you want to contribute to Immutable X documentation?
See our [docs repo](https://github.com/immutable/imx-docs/) and [contribution guidelines](https://github.com/immutable/imx-docs/blob/main/CONTRIBUTING.md).

## Getting started with contributing

There are two main ways to contribute:
1. [Opening an issue](#opening-an-issue)
2. [Sending a pull request with changes](#send-a-pull-request-with-changes)

## Opening an issue
Github issues are a great way to discuss and coordinate new features, changes, bugs, etc. You can create an issue [here](https://github.com/immutable/imx-core-sdk-golang/issues).

You should create an issue to open discussions:
* Before starting development on a significant pull request - use Github issues as a way to collaboratively propose, design, and discuss any development work you plan to do.
* Proposing new features or ideas - if you just have an idea or feature you'd like to see, drop the details into a Github issue for us to keep track of. Be specific with any requests, and include descriptions of what you would like to be added and what problems it would solve.
* Reporting a bug or security vulnerability - as above, opening an issue helps us keep track of and prioritize any bugs or vulnerabilities. Make sure to include all the relevant information, e.g. code snippets, error messages, stack traces, environment, reproduction steps, etc, to help us effectively diagnose and address the issue.

There is usually no need to create an issue when:
* An issue or open PR for the topic already exists - please continue or reopen the existing thread to keep all the discussion in one place.
* The PR is relatively small and obvious (e.g. small fixes, typos, etc) - if the changes do not require a discussion, simply filling out the PR description template will suffice.
* Requesting development help - if you run into any issues and require help setting up or using this SDK, feel free to reach out on [Discord](https://discord.com/invite/immutablex) in our #dev-discussion channel where we or the rest of the dev community will be happy to answer questions, or post a message on our [forum](https://forum.immutable.com/).
* Providing comments or general feedback - we generally use Github issues for actionable feedback or tasks (e.g. feature requests or bugs), but we welcome all forms of feedback. We would love to see you share your thoughts and comments in our Discord channel and engage with the rest of the community.

Issue titles should follow the below pattern:
```
update: Description # if an update is required for a feature
bug: Description # if there is a bug in a particular feature
suggestion: Description # if you want to suggest a better way to implement a feature
```

## Send a pull request with changes

To ensure high quality contributions, please follow the guidelines below:
* **Size:** Please keep pull requests small where possible (aim for 200-400 lines of code). Immutable prefers small pull requests to effectively review them.
* **Status checks:** Verify all status checks are passing. If a status check is failing, and you believe that the failure is unrelated to your change, please leave a comment on the pull request explaining why you believe the failure is unrelated. A maintainer will re-run the status check for you. If we conclude that the failure was a false positive, then we will open an issue to track that problem with our status check suite.
* Please fill out all fields in the [PR template](https://github.com/immutable/imx-core-sdk-golang/blob/main/.github/pull_request_template.md).

#### 1. Fork the [Immutable Golang Core SDK repository](https://github.com/immutable/imx-core-sdk-golang)
> To learn how to fork a repository, check out [GitHub's documentation on forking a repository](https://help.github.com/en/articles/fork-a-repo)


#### 2. Clone the forked repository:
```shell
git clone git@github.com:[your_github_handle]/imx-core-sdk-golang
```

#### 3. Navigate to the cloned folder:
```shell
cd imx-core-sdk-golang
```

#### 4. Link your cloned repository to the upstream repository:
```shell
git remote add upstream git@github.com:immutable/imx-core-sdk-golang.git
```

> For further help, see [GitHub Docs: Configuring a remote for a fork](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/configuring-a-remote-for-a-fork).

#### 5. If you have already cloned the imx-core-sdk-golang repository, ensure that your fork is up to date with with the latest upstream changes. 
> For further help, refer to [GitHub Docs: Syncing a fork](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork).
```shell
git checkout main
git fetch upstream
git merge upstream/main
```

#### 6. Create a new branch for your changes
```shell
git checkout -b [new_branch_name]
```

#### 7. Commit your changes while ensuring you follow our Git guidelines
```shell
git commit -m "feat: brief description of changes [Fixes #420]"
```

In the commit message, please reference the issue it resolves. For help, see [GitHub Docs: Linking a pull request to an issue using a keyword](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue).

Commit messages should follow the below pattern:
```feat: Description # if a new feature is added
fix: Description # if a bug is fixed
refactor: Description # if code is refactored
docs: Description # if documentation is added
lint: Description # if a lint issue is fixed
```

#### 8. Push your changes
```shell
git push
```

#### 9. The community maintainers / Immutable X team will review your PR

One of the community maintainers or the IMX team will review your PR and either accept it as it is, or add their comments as part of their review. Accepted PRs will get approved and merged into the `main` branch.

