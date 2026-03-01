# servaway

An HTTP server that returns 503 Service Unavailable on behalf of a service under maintenance.

## Overview

Responds to all requests with HTTP 503 Service Unavailable and a `Retry-After` header as a maintenance page.

## Usage

```
$ servaway
```

Listens on port 8080 by default.

### Options

```
$ servaway -port 3000 -retry-after 1800
```

| Flag | Description | Default |
| --- | --- | --- |
| `-port` | Listening port | `8080` |
| `-retry-after` | `Retry-After` header value in seconds | `3600` |

## Verification

```
$ curl -i http://localhost:8080/
HTTP/1.1 503 Service Unavailable
Content-Type: text/html; charset=utf-8
Retry-After: 3600
```

## Use Case

Switch the upstream of a reverse proxy such as Nginx to servaway during maintenance. After maintenance is complete, switch the upstream back to the original service.

```
[User] -> [Nginx] -> [Service]         (normal)
[User] -> [Nginx] -> [servaway]        (maintenance)
```

## License

This project is licensed under the [MIT License](./LICENSE).
