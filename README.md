# dbgproxy
The HTTP reverse proxy that logs everything in a single line

## How to run
```
docker run -it --net=host -e TARGET_PORT=8080 ghcr.io/pmh-only/dbgproxy
```

## Example log for nginx default page (beautified)
*Actual logs are single-lined*

```json
{
  "a": "GET / 200 1ms 615B 127.0.0.1 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:130.0) Gecko/20100101 Firefox/130.0'",
  "request": {
    "request_time": "2024-09-30 01:48:01.556622904 +0000 UTC",
    "request_method": "GET",
    "request_path": "/",
    "request_ip": "127.0.0.1",
    "request_headers": [
      {
        "header_key": "Host",
        "header_value": "localhost:3000"
      },
      {
        "header_key": "User-Agent",
        "header_value": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:130.0) Gecko/20100101 Firefox/130.0"
      },
      {
        "header_key": "Accept",
        "header_value": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/jxl,image/webp,image/png,image/svg+xml,*/*;q=0.8"
      },
      {
        "header_key": "Accept-Language",
        "header_value": "en-US,en;q=0.5"
      },
      {
        "header_key": "Accept-Encoding",
        "header_value": "gzip, deflate, br, zstd"
      },
      {
        "header_key": "Dnt",
        "header_value": "1"
      },
      {
        "header_key": "Sec-Gpc",
        "header_value": "1"
      },
      {
        "header_key": "Connection",
        "header_value": "keep-alive"
      },
      {
        "header_key": "Upgrade-Insecure-Requests",
        "header_value": "1"
      },
      {
        "header_key": "Sec-Fetch-Dest",
        "header_value": "document"
      },
      {
        "header_key": "Sec-Fetch-Mode",
        "header_value": "navigate"
      },
      {
        "header_key": "Sec-Fetch-Site",
        "header_value": "cross-site"
      },
      {
        "header_key": "Priority",
        "header_value": "u=0, i"
      },
      {
        "header_key": "Pragma",
        "header_value": "no-cache"
      },
      {
        "header_key": "Cache-Control",
        "header_value": "no-cache"
      }
    ],
    "request_body": "",
    "request_body_size": "0B"
  },
  "response": {
    "response_time": "2024-09-30 01:48:01.55710378 +0000 UTC",
    "response_code": 200,
    "response_headers": [
      {
        "header_key": "Content-Length",
        "header_value": "615"
      },
      {
        "header_key": "Content-Type",
        "header_value": "text/html"
      },
      {
        "header_key": "Server",
        "header_value": "nginx/1.27.1"
      },
      {
        "header_key": "Date",
        "header_value": "Mon, 30 Sep 2024 01:48:01 GMT"
      },
      {
        "header_key": "Last-Modified",
        "header_value": "Mon, 12 Aug 2024 14:21:01 GMT"
      },
      {
        "header_key": "Connection",
        "header_value": "keep-alive"
      },
      {
        "header_key": "Etag",
        "header_value": "\"66ba1a4d-267\""
      },
      {
        "header_key": "Accept-Ranges",
        "header_value": "bytes"
      }
    ],
    "response_body": "<!DOCTYPE html>\n<html>\n<head>\n<title>Welcome to nginx!</title>\n<style>\nhtml { color-scheme: light dark; }\nbody { width: 35em; margin: 0 auto;\nfont-family: Tahoma, Verdana, Arial, sans-serif; }\n</style>\n</head>\n<body>\n<h1>Welcome to nginx!</h1>\n<p>If you see this page, the nginx web server is successfully installed and\nworking. Further configuration is required.</p>\n\n<p>For online documentation and support please refer to\n<a href=\"http://nginx.org/\">nginx.org</a>.<br/>\nCommercial support is available at\n<a href=\"http://nginx.com/\">nginx.com</a>.</p>\n\n<p><em>Thank you for using nginx.</em></p>\n</body>\n</html>\n",
    "response_body_size": "615B"
  },
  "latency_ms": 1
}
```
