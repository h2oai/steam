Vagrant with YARN and PostgreSQL
====

This vagrant creates a CentOS 6.8 box with YARN and PostgreSQL set ready for SteamY usage. There are several final setup requirements before use.

1. Run the vagrant setup procedure.
    ```
    vagrant up
    vagrant ssh
    ```

2. Build the go project. (This needs to be done in a linux machine, cannot cross compile!)
    ```
    cd ./Go/src/github.com/steamY
    go build
    ```
    
3. In order to utilize yarn with steam, steam needs to be started as root. (The scoring service can be run using the default, no need for root.) 
    ```
    sudo env PATH=$PATH:/usr/local/hadoop/bin \
    ./steam serve master \
    --superuser-name=superuser \
    --superuser-password=superuser
    ```



