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

    Get a list of identities and roles with access to an entity
    $ steam get identities --for-entity \
        --entity-type=? \
        --entity-id=?

```
steam get identities [?]
```

### Options

```
      --entity-id=0: An entity ID.
      --entity-type=0: An entity type ID.
      --for-entity[=false]: Get a list of identities and roles with access to an entity
      --for-role[=false]: List identities for a role
      --for-workgroup[=false]: List identities for a workgroup
      --limit=0: The maximum returned objects.
      --offset=0: An offset to start the search on.
      --role-id=0: Integer ID of a role in Steam.
      --workgroup-id=0: Integer ID of a workgroup in Steam.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

