## steam stop cluster

Stop a cluster.

### Synopsis


Stop a cluster.

Examples:

    $ steam stop cluster 42

```
steam stop cluster [clusterId]
```

### Options

```
      --force[=false]: Force-kill all H2O instances in the cluster
      --kerberos[=true]: Set false on systems with no kerberos authentication.
      --keytab="": The name of the keytab file to use
      --username="": The valid kerberos username.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam stop](steam_stop.md)	 - Stop the specified resource.

