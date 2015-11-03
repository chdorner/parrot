VERSION := 0.1.0
LDFLAGS := -ldflags \
	"-X main.Version=$(VERSION)"

.PHONY: test deps

default:
	GO15VENDOREXPERIMENT=1 go build -v -o bin/parrot $(LDFLAGS)

deps:
	gvt rebuild

test:
	go list ./... | grep -v "vendor\/" | GO15VENDOREXPERIMENT=1 xargs go test
