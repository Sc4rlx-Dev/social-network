GO ?= go

.PHONY: tidy
tidy:
	$(GO) mod tidy