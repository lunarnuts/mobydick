#!/bin/bash

diff <(go run cmd/app/main.go $1) <(./mobydick.sh $1)
