# ---- Go-specific helpers ----------------------------------------------------
GOFILES     := $(shell find . -name '*.go' -type f)
GOMOD       := go.mod
GOSUM       := go.sum

DIST        := tuner1
BINARY      := dist/$(DIST)

# ---- Installation helpers (change if necessary) -----------------------------
INSTALL_DIR := /usr/local/bin/

# ---- Default target ---------------------------------------------------------
all: $(BINARY)

# ---- Build rule -------------------------------------------------------------
$(BINARY): $(GOFILES) $(GOMOD) $(GOSUM) config/standards.txt
	mkdir -p dist/
	go build -o $@ .

# ---- Install / upgrade / uninstall / clean ----------------------------------
install: upgrade
	set -x; \
	mkdir -p $(HOME)/.config/tuner1/; \
	OS=$$(uname -s | tr '[:upper:]' '[:lower:]'); \
	CONFIG_DIR=$$HOME/.config; \
	if [ "$$OS" = darwin ]; then \
		CONFIG_DIR="$$HOME/Library/Application Support"; \
	fi; \
	mkdir -p "$$CONFIG_DIR/tuner1/"; \
	cp -n config/standards.txt "$$CONFIG_DIR/tuner1/standards.txt"

upgrade: $(BINARY)
	sudo cp $(BINARY) /usr/local/bin/

uninstall:
	sudo rm -rf /usr/local/bin/$(DIST)

clean:
	rm -rf dist
