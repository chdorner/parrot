# parrot

    Parrot: Dexter's a cookie!
    Dexter: I am not a cookie!
    Parrot: Dexter's a cookie!
    Dexter: Am not!
    Parrot: Are too, cookie! COOKIE!
    Dexter: Good riddance! That has to be my worst invention yet!

Parrot is a small HTTP server that replies with the URL path you call it with.

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
