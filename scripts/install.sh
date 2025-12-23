#!/usr/bin/env bash
set -euo pipefail

REPO="lxsavage/tuner1"
BINARY="tuner1"

# ---- 1. Map OS / Arch to release naming ----
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
  linux)  OS="linux" ;;
  darwin) OS="macos" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

ASSET="${BINARY}-${OS}-${ARCH}"

# ---- 2. Pick install directory ----
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
mkdir -p "$INSTALL_DIR"

# ---- 3. Download URL for latest release ----
URL=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" \
        | grep "browser_download_url.*$ASSET" \
        | head -n 1 \
        | cut -d '"' -f 4)

if [[ -z "$URL" ]]; then
  echo "Release asset $ASSET not found in latest release"
  exit 1
fi

# ---- 4. Download & install ----
TMP=$(mktemp -d)
trap "rm -rf $TMP" EXIT
curl -sSL "$URL" -o "$TMP/$ASSET"
chmod +x "$TMP/$ASSET"
mv "$TMP/$ASSET" "$INSTALL_DIR/$BINARY"

# ---- 5. Pull the latest standards.txt into the tuner1 config dir if not already there
CONFIG_DIR=$HOME/.config
if [[ $OS -eq darwin ]]; then
  CONFIG_DIR="$HOME/Library/Application Support"
else
  mkdir -p $CONFIG_DIR
fi

mkdir -p $CONFIG_DIR/tuner1
STANDARDS_FILE="$CONFIG_DIR/tuner1/standards.txt"

if [ ! -e "$STANDARDS_FILE" ]; then
  curl -sSL \
    "https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/config/standards.txt" \
    -o "$STANDARDS_FILE" --no-clobber
else
    echo "Standards file already exists locally. Skipping download."
fi

# ---- 5. Final message ----
echo "$BINARY installed to $INSTALL_DIR/$BINARY and standards.txt to $STANDARDS_FILE"
if ! command -v "$BINARY" >/dev/null 2>&1; then
  echo "Add $INSTALL_DIR to your PATH, e.g.:"
  echo "  export PATH=\"$INSTALL_DIR:\$PATH\""
fi
