NAME := qv-tones-player
GO   := /opt/go

all: clean build

build:
	go build -o $(NAME).x64.linux
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=windows GOARCH=386   go build -o $(NAME).win32.exe

clean:
	rm -f *.linux *.exe
