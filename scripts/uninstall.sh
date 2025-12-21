#!/bin/sh
#!/usr/bin/env bash
set -euo pipefail

BINARY="tuner1"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

BINARY_PATH="$INSTALL_DIR/$BINARY"

if [[ ! -x "$BINARY_PATH" ]]; then
  echo "$BINARY not found in $INSTALL_DIR; nothing to uninstall."
  exit 0
fi

read -p "Remove $BINARY_PATH? [y/N] " -n 1 -r

echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
  rm -f "$BINARY_PATH"
  echo "$BINARY removed."
  printf "To remove the standards file, run:\n$ rm $HOME/.config/tuner1/standards.txt\n"
else
  echo "Uninstall cancelled."
fi
