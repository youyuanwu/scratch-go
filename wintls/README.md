This is an attempt to use cert in windows cert store to run a https server.

It seems that tls handshake fails:
server:
```
http: TLS handshake error from [::1]:xxxxxx: remote error: tls: error decrypting message
```

client:
```
Get "https://localhost:12345": tls: invalid signature by the server certificate: crypto/rsa: verification error
```