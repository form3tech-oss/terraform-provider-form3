#!/usr/bin/env bash

# We only need to run goreleaser to attach the binaries
# to the github release if this is a build triggered
# by a tag
if [ -n "$TRAVIS_TAG" ]; then
    curl -sL http://git.io/goreleaser | bash
else
    echo "not running goreleaser for this build as TRAVIS_TAG is empty"
fi
