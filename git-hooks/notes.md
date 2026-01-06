# Installing hooks

## Prerequisite

golangci-lint is required to be installed for pre-commit hook to work. For instructions regarding installing golangci-lint, see https://golangci-lint.run/usage/install/#local-installation

## Installing Pre Commit Hook
After cloning a git repo to your machine git hooks aren’t installed automatically. You’ll need to store them at your repository and reinstall each time you change their code or clone the repository to a new place.

cp pre-commit ../.git/hooks/pre-commit
