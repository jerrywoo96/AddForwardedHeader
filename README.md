# Add Forwarded Header

Construct Forwarded Header from X-Forwarded-For, X-Forwarded-Host and X-Forwarded-Proto incoming request headers, and appends them to outgoing request headers.

See [https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded)

Traefik Plugin: [https://plugins.traefik.io/plugins/62cfd4129279ff6d9dd027a9/add-forwarded-header](https://plugins.traefik.io/plugins/62cfd4129279ff6d9dd027a9/add-forwarded-header)

GitHub: [https://github.com/jerrywoo96/AddForwardedHeader](https://github.com/jerrywoo96/AddForwardedHeader)

## Configuration

### Static (traefik.yml)

```yaml
experimental:
  plugins:
    AddForwardedHeader:
      moduleName: github.com/jerrywoo96/AddForwardedHeader
      version: v1.0.1
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
