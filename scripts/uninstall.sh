#!/usr/bin/env bash

BINARY="tuner1"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

BINARY_PATH="$INSTALL_DIR/$BINARY"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')

if [[ ! -x "$BINARY_PATH" ]]; then
  echo "$BINARY not found in $INSTALL_DIR; nothing to uninstall."
  exit 0
fi

CONFIG_DIR=$HOME/.config
if [[ $OS -eq darwin ]]; then
  CONFIG_DIR="$HOME/Library/Application Support"
fi

rm -f "$BINARY_PATH"
echo "$BINARY removed."
printf "To remove the standards file, run:\n$ rm $CONFIG_DIR/tuner1/standards.txt\n"
