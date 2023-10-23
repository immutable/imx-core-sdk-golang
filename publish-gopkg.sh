#!/bin/zsh

VERSION_STR="v1.0.1"

git tag ${VERSION_STR}
git push origin ${VERSION_STR}

GOPROXY=proxy.golang.org go list -m github.com/immutable/imx-core-sdk-golang@${VERSION_STR}

exit 0
