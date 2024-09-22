This is an attempt to use cert in windows cert store to run a https server.

```ps1
New-SelfSignedCertificate -Type Custom -Subject "CN=Youyuan Wu Test" `
    -KeyExportPolicy NonExportable -KeyUsage DigitalSignature -KeyAlgorithm RSA -KeyLength 2048 -CertStoreLocation "Cert:\CurrentUser\My" `
    -Provider "Microsoft Platform Crypto Provider"
```