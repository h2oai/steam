## steam build model

Build Model

### Synopsis


Build Model

Examples:

    Build a model
    $ steam build model \
        --cluster-id=? \
        --dataset-id=? \
        --algorithm=?

    Build an AutoML model
    $ steam build model --auto \
        --cluster-id=? \
        --dataset=? \
        --target-name=? \
        --max-run-time=?

```
steam build model [?]
```

### Options

```
      --algorithm="": No description available
      --auto[=false]: Build an AutoML model
      --cluster-id=0: No description available
      --dataset="": No description available
      --dataset-id=0: No description available
      --max-run-time=0: No description available
      --target-name="": No description available
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam build](steam_build.md)	 - Build entities

