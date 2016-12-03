#!/usr/bin/env bash

go install -ldflags "-X main.buildTime=`date -u +%Y.%m.%d-%Hh%Mm%S`"


