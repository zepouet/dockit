#!/usr/bin/env bash

go build -ldflags "-X main.buildTime=`date -u +%Y.%m.%d-%Hh%Mm%S`" -o dockit main.go
./dockit version


