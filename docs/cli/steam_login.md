## steam login

Sign in to a Steam server.

### Synopsis


Sign in to a Steam server.

Examples:

	$ steam login 192.168.42.42:9000 \
			--username=arthur
			--password=beeblebrox

```
steam login [address:port] --username=[username] --password=[password]
```

### Options

```
      --authentication="basic": Authentication method
      --password="": Login password
      --secure[=false]: Enable TLS
      --username="": Login username
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam](steam.md)	 - steam v build : Command Line Interface to Steam

