## steam get roles

Get Roles

### Synopsis


Get Roles

Examples:

    List roles
    $ steam get roles \
        --offset=? \
        --limit=?

    List roles for an identity
    $ steam get roles --for-identity \
        --identity-id=?

```
steam get roles [?]
```

### Options

```
      --for-identity[=false]: List roles for an identity
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

