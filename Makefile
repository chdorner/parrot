VERSION := 0.2.0
VENDORENV := GO15VENDOREXPERIMENT=1
LDFLAGS := -ldflags \
	"-X main.Version=$(VERSION)"

.PHONY: test deps

default:
	$(VENDORENV) go build -v -o bin/parrot $(LDFLAGS)

deps:
	gvt rebuild

test:
	go list ./... | grep -v "vendor\/" | $(VENDORENV) xargs go test

release:
	$(VENDORENV) GOOS=darwin GOARCH=amd64 go build -o pkg/parrot-$(VERSION)_amd64_darwin
	$(VENDORENV) GOOS=linux GOARCH=amd64 go build -o pkg/parrot-$(VERSION)_amd64_linux
	$(VENDORENV) GOOS=freebsd GOARCH=amd64 go build -o pkg/parrot-$(VERSION)_amd64_freebsd
	$(VENDORENV) GOOS=linux GOARCH=386 go build -o pkg/parrot-$(VERSION)_386_linux
	$(VENDORENV) GOOS=freebsd GOARCH=386 go build -o pkg/parrot-$(VERSION)_386_freebsd
