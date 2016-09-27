## steam get datasets

Get Datasets

### Synopsis


Get Datasets

Examples:

    List datasets
    $ steam get datasets \
        --datasource-id=? \
        --offset=? \
        --limit=?

    Get a list of datasets on a cluster
    $ steam get datasets --from-cluster \
        --cluster-id=?

```
steam get datasets [?]
```

### Options

```
      --cluster-id=0: No description available
      --datasource-id=0: No description available
      --from-cluster[=false]: Get a list of datasets on a cluster
      --limit=0: No description available
      --offset=0: No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

