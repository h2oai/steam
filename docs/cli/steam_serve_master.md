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
      --db-connection-timeout="": Database connection timeout (optional)
      --db-host="": Database host (optional, defaults to localhost
      --db-name="steam": Database name to use for application data storage (required)
      --db-password="": Database password (optional)
      --db-port="": Database port (optional, defaults to 5432)
      --db-ssl-cert-path="": Database connection SSL certificate path (optional)
      --db-ssl-key-path="": Database connection SSL key path (optional)
      --db-ssl-mode="disable": Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'
      --db-ssl-root-cert-path="": Database connection SSL root certificate path (optional)
      --db-username="steam": Database username (required)
      --profile[=false]: Enable Go profiler
      --scoring-service-address="": Address to start scoring services on ("<ip>")
      --scoring-service-port-range="1025:65535": Specified port range to create scoring services on. ("<from>:<to>")
      --admin-name="": Set admin username (required for first-time-use only)
      --admin-password="": Set admin password (required for first-time-use only)
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
* [steam serve](steam_serve.md)	 - Launch a new service.

