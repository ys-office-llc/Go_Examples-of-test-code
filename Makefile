# Makefile

test:
	go test -coverprofile=coverage.out ./...

cover:
	go tool cover -html=coverage.out -o coverage.html

open:
	xdg-open coverage.html
