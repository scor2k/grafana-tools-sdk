#go generate ./...

git config core.hooksPath .githooks

BUILD_HASH=$(git rev-parse --verify HEAD)
BUILD_TIME=$(git show -s --format=%ci)

go install -tags netgo,osusergo -v -ldflags="-X 'github.com/enuan/m3-go/core.buildHash=$BUILD_HASH' -X 'github.com/enuan/m3-go/core.buildTime=$BUILD_TIME'" ./...
