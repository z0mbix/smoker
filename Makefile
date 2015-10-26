PREFIX := /usr/local

build:
	go build smoker.go

run:
	go run smoker.go

clean:
	rm smoker

install:
	install -m 0755 smoker $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin/smoker
