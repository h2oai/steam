## steam find models

Find Models

### Synopsis


Find Models

Examples:

    Get a count models in a project
    $ steam find models --count \
        --project-id=?

    List binomial models
    $ steam find models --binomial \
        --project-id=? \
        --name-part=? \
        --sort-by=? \
        --ascending=? \
        --offset=? \
        --limit=?

    List multinomial models
    $ steam find models --multinomial \
        --project-id=? \
        --name-part=? \
        --sort-by=? \
        --ascending=? \
        --offset=? \
        --limit=?

    List regression models
    $ steam find models --regression \
        --project-id=? \
        --name-part=? \
        --sort-by=? \
        --ascending=? \
        --offset=? \
        --limit=?

```
steam find models [?]
```

### Options

```
      --ascending[=false]: No description available
      --binomial[=false]: List binomial models
      --count[=false]: Get a count models in a project
      --limit=0: No description available
      --multinomial[=false]: List multinomial models
      --name-part="": No description available
      --offset=0: No description available
      --project-id=0: No description available
      --regression[=false]: List regression models
      --sort-by="": No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam find](steam_find.md)	 - Find entities

