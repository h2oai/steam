
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

if go is already installed, check go version: 

```
$ go version
```

if the go version displayed is below `1.7.0`, follow the instructions at golang.org to:
- [uninstall the old version of go](https://golang.org/doc/install#uninstall)
- [install the latest version of go](https://golang.org/doc/install)

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
$ cd $GOPATH/src/github.com/h2oai/
$ git clone https://github.com/h2oai/steam
$ cd steam
$ make
```

Note: If you do not follow above path to clone steam source, you will get build errors. golang is very specific about where the source code is cloned.

**Step 5**

After your build is successful you will see steam binary is available. Now initialize the (sqlite) database required for steam. To initialize the sqlite database:

```
make db
```

Now you are ready to start the server. 

**Step 6**
Start the server as below

```
./steam serve master --superuser-name USER_NAME --superuser-password PASSWORD
```






