#!/bin/bash
cd "$(dirname "$0")"
go install github.com/dennwc/flatpak-go-mod@latest
flatpak-go-mod $1
