## steam get package

Get Package

### Synopsis


Get Package

Examples:

    List directories in a project package
    $ steam get package --directories \
        --project-id=? \
        --package-name=? \
        --relative-path=?

    List files in a project package
    $ steam get package --files \
        --project-id=? \
        --package-name=? \
        --relative-path=?

```
steam get package [?]
```

### Options

```
      --directories[=false]: List directories in a project package
      --files[=false]: List files in a project package
      --package-name="": No description available
      --project-id=0: No description available
      --relative-path="": No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam get](steam_get.md)	 - Get entities

