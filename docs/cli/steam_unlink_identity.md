## steam unlink identity

Unlink Identity

### Synopsis


Unlink Identity

Examples:

    Unlink an identity from a workgroup
    $ steam unlink identity --from-workgroup \
        --identity-id=? \
        --workgroup-id=?

    Unlink an identity from a role
    $ steam unlink identity --from-role \
        --identity-id=? \
        --role-id=?

```
steam unlink identity [?]
```

### Options

```
      --from-role[=false]: Unlink an identity from a role
      --from-workgroup[=false]: Unlink an identity from a workgroup
      --identity-id=0: Integer ID of an identity in Steam.
      --role-id=0: Integer ID of a role in Steam.
      --workgroup-id=0: Integer ID of a workgroup in Steam.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam unlink](steam_unlink.md)	 - Unlink entities

