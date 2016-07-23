## steam get identities

Get Identities

### Synopsis


Get Identities

Examples:

    List identities
    $ steam get identities \
        --offset=? \
        --limit=?

    List identities for a workgroup
    $ steam get identities --for-workgroup \
        --workgroup-id=?

    List identities for a role
    $ steam get identities --for-role \
        --role-id=?

```
steam get identities [?]
```

### Options

```
      --for-role[=false]: List identities for a role
      --for-workgroup[=false]: List identities for a workgroup
      --limit=0: No description available
      --offset=0: No description available
      --role-id=0: No description available
      --workgroup-id=0: No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

