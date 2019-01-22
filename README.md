# DNSRepeater (Server) [![Build Status](https://travis-ci.org/FergusInLondon/DNSRepeat.svg?branch=master)](https://travis-ci.org/FergusInLondon/DNSRepeat)

DNSRepeater is a small golang microservice that provides a *very simple* interface for providing DNS lookups over HTTP(S) - it's intended to support simple web browsing only. (for situations where DNS tampering may be employed as a form of censorship)

It was wrote over the period of one evening, and as such may miss certain edge cases - although it has a comprehensive set of tests.

## Rationale

During a conversation on `/r/sysadmin` about the censorship of Venezuelan internet access under the current regime, one of key points was that the censorship relies on DNS tampering.

Whilst it's possible to adjust DNS settings on a localhost - i.e by utilising OpenDNS or Google Public DNS - this can be mitigated at an ISP level by blacklisting those services, or filtering out DNS traffic that doesn't use their own servers.

A potential workaround for this is to proxy DNS requests over an alternative transport protocol - the simplest being HTTP, which can be secured via HTTPS and appears as normal web traffic.

To utilise this the client would needs to be able to resolve DNS requests locally and proxy the requests to the HTTP service.. but more on that later.

## Development

This has been compiled without any issues using Go 1.11 on an Arch Linux derivative.

## Testing

There's fairly comprehensive test coverage that *should* cover a multitude of edge cases.

    ➜  DNSRepeat git:(master) ✗ go test -cover
    PASS
    coverage: 84.1% of statements
    ok      _/home/fergus/Code/DNSRepeat    0.008s


## Deployment

Deployment is trivial via Docker. (See `Dockerfile`)

## Examples

A note: *yes, we're using the request body on a `GET` request*. This is a bit of an anti-pattern, but it's certainly valid.

#### Resolve the IP address for `github.com`

**Request**

    GET: /
    {
        "hostname": "github.com"
    }

**Response**:

    Status: 200,
    Content-Type: application/json
    {
        "hostname": "github.com"
        "address":  "140.82.118.3"
    }

#### Resolve the IP address for `gist.github.com`

**Request**:

    GET: /
    {
        "hostname": "gist.github.com"
    }

**Response**:

    Status: 200,
    Content-Type: application/json
    {
        "hostname": "gist.github.com"
        "address":  "192.30.253.119"
    }

## Todo

- Optional Debug Logging (Anonymised)
- Cache persistence (dump upon signal, read upon init)
- Configuration via Environment Variables