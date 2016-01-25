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
	$(VENDORENV) GOOS=darwin GOARCH=amd64 go build -o pkg/parrot-$(VERSION)_amd64_darwin $(LDFLAGS)
	$(VENDORENV) GOOS=linux GOARCH=amd64 go build -o pkg/parrot-$(VERSION)_amd64_linux $(LDFLAGS)
	$(VENDORENV) GOOS=freebsd GOARCH=amd64 go build -o pkg/parrot-$(VERSION)_amd64_freebsd $(LDFLAGS)
	$(VENDORENV) GOOS=linux GOARCH=386 go build -o pkg/parrot-$(VERSION)_386_linux $(LDFLAGS)
	$(VENDORENV) GOOS=freebsd GOARCH=386 go build -o pkg/parrot-$(VERSION)_386_freebsd $(LDFLAGS)
