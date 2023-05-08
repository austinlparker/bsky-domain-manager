#!/bin/bash

set -e

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

if [[ "$GOOS" == "windows" ]]; then
  SUFFIX=".exe"
else
  SUFFIX=""
fi

if [[ ! -d /usr/local/bin ]]; then
  sudo mkdir -p /usr/local/bin
fi

sudo cp skeeter-$GOOS-$GOARCH$SUFFIX /usr/local/bin/skeeter
sudo chmod +x /usr/local/bin/skeeter
