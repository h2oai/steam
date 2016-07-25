## steam get cluster

Get Cluster

### Synopsis


Get Cluster

Examples:

    Get cluster details
    $ steam get cluster \
        --cluster-id=?

    Get cluster details (Yarn only)
    $ steam get cluster --on-yarn \
        --cluster-id=?

    Get cluster status
    $ steam get cluster --status \
        --cluster-id=?

```
steam get cluster [?]
```

### Options

```
      --cluster-id=0: No description available
      --on-yarn[=false]: Get cluster details (Yarn only)
      --status[=false]: Get cluster status
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

