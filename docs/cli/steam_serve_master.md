## steam serve master

Launch the Steam master.

### Synopsis


Launch the Steam master.

Examples:

    $ steam serve master

```
steam serve master
```

### Options

```
      --authentication-provider="basic": Authentication mechanismfor client logins (one of "basic" or "digest")
      --cluster-proxy-address=":9001": Cluster proxy address ("<ip>:<port>" or ":<port>")
      --compilation-service-address=":8080": Model compilation service address ("<ip>:<port>")
      --db-name="steam": Database name to use for application data storage
      --db-ssl-mode="disable": Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'
      --db-username="steam": Database username to connect as
      --profile[=false]: Enable Go profiler
      --scoring-service-address="": Address to start scoring services on ("<ip>")
      --superuser-name="": Set superuser username (required for first-time-use only)
      --superuser-password="": Set superuser password (required for first-time-use only)
      --web-address=":9000": Web server address ("<ip>:<port>" or ":<port>").
      --web-tls-cert-path="": Web server TLS certificate file path (optional).
      --web-tls-key-path="": Web server TLS key file path (optional).
      --working-directory="var/master": Working directory for application files.
      --yarn-enable-kerberos[=false]: Enable Kerberos authentication. Requires username and keytab.
      --yarn-keytab="": Keytab file to be used with Kerberos authentication
      --yarn-username="": Username to enable Kerberos
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam serve](steam_serve.md)	 - Lauch a new service.

