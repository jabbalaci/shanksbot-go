cat:
	cat Makefile

bot:
	go run ./cmd/shanksbot

gen:
	go run ./cmd/generator

#############################################################################

test:
	go test ./... -v
