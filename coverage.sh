#!/bin/bash

go test -coverprofile=cover.out ./poet && go tool cover -html=cover.out -o coverage.html && echo "Coverage report created in $PWD/coverage.html"

