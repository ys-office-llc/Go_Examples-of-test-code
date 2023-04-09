# Makefile

test:
	go test -v -coverprofile=coverage.out ./...

cover:
	go tool cover -html=coverage.out -o coverage.html

open:
	xdg-open coverage.html
