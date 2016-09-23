## steam get history

Get History

### Synopsis


Get History

Examples:

    List audit trail records for an entity
    $ steam get history \
        --entity-type-id=? \
        --entity-id=? \
        --offset=? \
        --limit=?

```
steam get history [?]
```

### Options

```
      --entity-id=0: Integer ID for an entity in Steam.
      --entity-type-id=0: Integer ID for the type of entity.
      --limit=0: The maximum returned objects.
      --offset=0: An offset to start the search on.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

