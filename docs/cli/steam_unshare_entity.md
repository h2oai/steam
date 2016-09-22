## steam unshare entity

Unshare Entity

### Synopsis


Unshare Entity

Examples:

    Unshare an entity
    $ steam unshare entity \
        --kind=? \
        --workgroup-id=? \
        --entity-type-id=? \
        --entity-id=?

```
steam unshare entity [?]
```

### Options

```
      --entity-id=0: Integer ID for an entity in Steam.
      --entity-type-id=0: Integer ID for the type of entity.
      --kind="": Type of permission. Can be view, edit, or own.
      --workgroup-id=0: Integer ID of a workgroup in Steam.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam unshare](steam_unshare.md)	 - Unshare entities

