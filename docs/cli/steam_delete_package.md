## steam delete package

Delete Package

### Synopsis


Delete Package

Examples:

    Delete a project package
    $ steam delete package \
        --project-id=? \
        --name=?

    Delete a directory in a project package
    $ steam delete package --directory \
        --project-id=? \
        --package-name=? \
        --relative-path=?

    Delete a file in a project package
    $ steam delete package --file \
        --project-id=? \
        --package-name=? \
        --relative-path=?

```
steam delete package [?]
```

### Options

```
      --directory[=false]: Delete a directory in a project package
      --file[=false]: Delete a file in a project package
      --name="": No description available
      --package-name="": No description available
      --project-id=0: No description available
      --relative-path="": No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam delete](steam_delete.md)	 - Delete entities

