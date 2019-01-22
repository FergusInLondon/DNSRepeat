# DNSRepeater (Server)

This is a simple little project that retrieves DNS records and sends them via a JSON response back to a client; in essence a crappy DNS proxy that can work over HTTP(S). In reality I'd like to expand upon it, docker it up, and see if it's a potential strategy to help fight against web censorship where DNS tampering is employed. (*How?* Essentially by using a local caching DNS server that's capable of resolving via HTTPS, but more on that later.)

Right now I'm a bit tired, as this was a "pre-bedtime" prototype more than anything - and I ended up getting a little sidetracked whilst working on it.

There's decent test coverage, a complete lack of comments and a complet lack of documentation; needless to say they're tomorrows problems.

It also has no logging; but whether that's a good thing or not considering it's use-case.. well, remains to be seen. A debug flag will no doubt be enabled later.

## Examples

Here's a simple example

Request 1:

    GET: /
    {
        "hostname": "github.com"
    }

Response:

    Status: 200,
    Content-Type: application/json
    {
        "hostname": "github.com"
        "address":  "140.82.118.3"
    }


Request 2:

    GET: /
    {
        "hostname": "gist.github.com"
    }

Response 2:

    Status: 200,
    Content-Type: application/json
    {
        "hostname": "gist.github.com"
        "address":  "192.30.253.119"
    }
