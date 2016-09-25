
# Development setup instructions

**Step 1**

Install Node.js (https://nodejs.org/en/)

###### Preparing NPM private registry
Create an ```.npmrc``` file in the root folder - sample ```.npmrc``` can be found in [this gist](https://gist.github.com/justinloyola/29229513843722390fda39752439a5e8).

The _auth value is a base64 encoding of Nexus Repository credentials. The credentials are available through Dev Ops.

```
echo -n 'username:password' | openssl base64
```

Once the private registry has been setup, you can publish to the private registry using ```npm publish```. Any repository that installs from the private registry will also have to have an ```.npmrc``` file in its root folder.

Next, simply ```npm start``` to start the frontend development HTTP server on port 3000.

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
