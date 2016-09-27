## steam get workgroups

Get Workgroups

### Synopsis


Get Workgroups

Examples:

    List workgroups
    $ steam get workgroups \
        --offset=? \
        --limit=?

    List workgroups for an identity
    $ steam get workgroups --for-identity \
        --identity-id=?

```
steam get workgroups [?]
```

### Options

```
      --for-identity[=false]: List workgroups for an identity
      --identity-id=0: Integer ID of an identity in Steam.
      --limit=0: The maximum returned objects.
      --offset=0: An offset to start the search on.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

