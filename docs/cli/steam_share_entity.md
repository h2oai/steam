## steam share entity

Share Entity

### Synopsis


Share Entity

Examples:

    Share an entity with a workgroup
    $ steam share entity \
        --kind=? \
        --workgroup-id=? \
        --entity-type-id=? \
        --entity-id=?

```
steam share entity [?]
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
* [steam share](steam_share.md)	 - Share entities

