#! /usr/bin/bash
set -xe

CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o ./build/trainer_editor.exe

