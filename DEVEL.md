
# Development setup instructions

**Step 1**

Install Node.js (https://nodejs.org/en/)

**Step 2**

Install Go from golang.org (use the OSX installer), and then:
```
$ which go
```

- If this points to /usr/local/go/bin/go, go to **Step 3**.
- If not, restart your terminal and try again.
- If not, add `export PATH=$PATH:/usr/local/go/bin` to your ~/.bash_profile and then `$ source ~/.bash_profile`.

**Step 3**

```
$ mkdir -p $HOME/go
```

Then, edit your ~/.bash_profile and add `export GOPATH=$HOME/go` and then `$ source ~/.bash_profile`.

Next, install this Go dependency management tool.

```
go get github.com/tools/godep
```

**Step 4**

Finally, clone this repo:

```
$ mkdir -p $GOPATH/src/github.com/h2oai
$ cd go/src/github.com/h2oai/
$ git clone https://github.com/h2oai/steam
$ cd steam
$ make
```
