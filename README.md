# parrot [![Build Status](https://travis-ci.org/chdorner/parrot.svg)](https://travis-ci.org/chdorner/parrot)

    Parrot: Dexter's a cookie!
    Dexter: I am not a cookie!
    Parrot: Dexter's a cookie!
    Dexter: Am not!
    Parrot: Are too, cookie! COOKIE!
    Dexter: Good riddance! That has to be my worst invention yet!

Parrot is a small HTTP server that replies with the URL path you call it with.

## Installation

If you already have Go >= 1.5 installed and set up your GOPATH:

```bash
$ go get -u github.com/chdorner/parrot
$ $GOPATH/bin/parrot -version
```

You can find binary distributions in the
[releases](https://github.com/chdorner/parrot/releases), install them like:

```bash
$ wget https://github.com/chdorner/parrot/releases/download/v0.1.0/parrot-0.1.0_amd64_linux
$ mv parrot-0.1.0_amd64_linux /usr/local/bin/parrot
$ parrot -version
```

Note that you might have to replace the download URL depending on your
operating system.

## Usage

```bash
# start the server
$ parrot -a :4242 &

# it knows about plain text
$ curl http://localhost:4242/i-am-a-parrot
/i-am-a-parrot

# it knows about json
$ curl http://localhost:4242/i-am-a-parrot.json
{"url":"/i-am-a-parrot.json"}

# it knows about xml
$ curl http://localhost:4242/i-am-a-parrot.xml
<parrot><url>/i-am-a-parrot.xml</url></parrot>

# it can also reply with http response codes to special requests
$ curl http://localhost:4242/_/201
201 - Created

$ curl http://localhost:4242/_/402.json
{"code":402,"text":"Payment Required"}

$ curl http://localhost:4242/_/505.xml
<status><code>418</code><text>I&#39;m a teapot</text></status>
```

## Notes on development

I'm using the new Go 1.5 vendor experiment, so Go 1.5 is required for building
parrot. Make sure `GO15VENDOREXPERIMENT=1` is set in your environment, or let
the `make` handle this for you. I'm also using
[gvt](https://github.com/FiloSottile/gvt) for handling the
fetching of the vendored packages.

```bash
# Get all dependencies
$ make deps

# Run the tests
$ make test

# Build from source
$ make

# Run
$ bin/parrot
```
