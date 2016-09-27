## steam link identity

Link Identity

### Synopsis


Link Identity

Examples:

    Link an identity with a workgroup
    $ steam link identity --with-workgroup \
        --identity-id=? \
        --workgroup-id=?

    Link an identity with a role
    $ steam link identity --with-role \
        --identity-id=? \
        --role-id=?

```
steam link identity [?]
```

### Options

```
      --identity-id=0: Integer ID of an identity in Steam.
      --role-id=0: Integer ID of a role in Steam.
      --with-role[=false]: Link an identity with a role
      --with-workgroup[=false]: Link an identity with a workgroup
      --workgroup-id=0: Integer ID of a workgroup in Steam.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam link](steam_link.md)	 - Link entities

