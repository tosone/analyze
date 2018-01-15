BuildStamp = main.BuildStamp=$(shell date '+%Y-%m-%d_%I:%M:%S%p')
GitHash    = main.GitHash=$(shell git rev-parse HEAD)
Version    = main.Version=$(shell git describe --abbrev=0 --tags)
Target     = worklyzer

UNAME_S    = $(shell uname -s)

GOOS       = windows
CC         = "x86_64-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp"
Subfix     = windows

ifeq ($(UNAME_S),Darwin)
	GOOS   = darwin
	CC     = clang
	Subfix = mac 
endif

ifeq ($(UNAME_S),Linux)
	GOOS       = linux
	CC         = gcc
	Subfix     = linux
endif

build: clean
	mkdir release
	CGO_ENABLED=1 CC=$(CC) GOOS=$(GOOS) GOARCH=amd64 go build -v -o release/${Target}-$(Subfix) -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" main.go

test: cleanTest
	./release/${Target}-$(Subfix) --config=config.yaml service

release:
	xgo -v --targets=linux/amd64,linux/arm-5,linux/arm-6,linux/arm-7,darwin/amd64,windows/amd64 -ldflags "-s -w -X ${BuildStamp} -X ${GitHash} -X ${Version}" github.com/tosone/worklyzer

authors:
	echo "Authors\n=======\n\nProject's contributors:\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

clean:
	-rm -rf release

cleanTest:
	-rm -rf *.log *.db