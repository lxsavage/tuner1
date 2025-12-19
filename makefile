all: dist/tuner1

dist/tuner1:
	mkdir -p dist/
	cp config/standards.txt dist/
	go get .
	go build -o dist/tuner1 .

install: dist/tuner1
	go get .
	cp dist/tuner1 /usr/local/bin/tuner1
	mkdir -p $$HOME/.config/tuner1/
	cp dist/standards.txt $$HOME/.config/tuner1/

upgrade: dist/tuner1
	rm /usr/local/bin/tuner1
	cp dist/tuner1 /usr/local/bin/tuner1

uninstall:
	rm -rf $$HOME/.config/tuner1/
	rm /usr/local/bin/tuner1

clean:
	rm -rf dist/
