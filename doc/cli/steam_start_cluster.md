## steam start cluster

Start a new cluster using the specified H2O package.

### Synopsis


Start a new cluster using the specified H2O package.

Examples:

Start a 4 node H2O 3.2.0.9 cluster

    $ steam start cluster42 2 --size=4

```
steam start cluster [clusterName] [engineId]
```

### Options

```
  -m, --memory="10g": The max amount of memory to use per node.
  -n, --size=1: The number of nodes to provision.
```

### Options inherited from parent commands

```
  -v, --verbose[=false]: verbose output
```

### SEE ALSO
* [steam start](steam_start.md)	 - Start a new resource.

