# AddForwardedHeader
Construct Forwarded Header from X-Forwarded-For, X-Forwarded-Host and X-Forwarded-Proto incoming request headers, and appends them to outgoing request headers.

## Configuration

### Static (traefik.yml)
```yaml
experimental:
  plugins:
    AddForwardedHeader:
      moduleName: github.com/jerrywoo96/AddForwardedHeader
      version: v1.0.0
```

### Dynamic
```yaml
http:
  middlewares:
    AddForwardedHeader:
      plugin:
        AddForwardedHeader:
          by: 10.0.0.1:443 # The IP that traefik uses to communicate with the destination server.
          forHeader: X-Real-Ip # Get the Client IP from this header.
```
