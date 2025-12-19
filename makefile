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
	mkdir -p $(HOME)/.config/tuner1/
	cp config/standards.txt $(HOME)/.config/tuner1/

upgrade: $(BINARY)
	cp $(BINARY) /usr/local/bin/

uninstall:
	rm -rf $(HOME)/.config/tuner1/
	rm -rf /user/local/bin/$(DIST)

clean:
	rm -rf dist
