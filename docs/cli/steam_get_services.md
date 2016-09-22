## steam get services

Get Services

### Synopsis


Get Services

Examples:

    List all services
    $ steam get services \
        --offset=? \
        --limit=?

    List services for a project
    $ steam get services --for-project \
        --project-id=? \
        --offset=? \
        --limit=?

    List services for a model
    $ steam get services --for-model \
        --model-id=? \
        --offset=? \
        --limit=?

```
steam get services [?]
```

### Options

```
      --for-model[=false]: List services for a model
      --for-project[=false]: List services for a project
      --limit=0: No description available
      --model-id=0: No description available
      --offset=0: No description available
      --project-id=0: No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

