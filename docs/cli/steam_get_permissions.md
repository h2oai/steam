## steam get permissions

Get Permissions

### Synopsis


Get Permissions

Examples:

    List permissions for a role
    $ steam get permissions --for-role \
        --role-id=?

    List permissions for an identity
    $ steam get permissions --for-identity \
        --identity-id=?

```
steam get permissions [?]
```

### Options

```
      --for-identity[=false]: List permissions for an identity
      --for-role[=false]: List permissions for a role
      --identity-id=0: No description available
      --role-id=0: No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

