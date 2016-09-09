Vagrant with YARN and PostgreSQL
====

This vagrant creates a CentOS 6.8 box with YARN and PostgreSQL set ready for SteamY usage. There are several final setup requirements before use.

1. Create linux steam. From outside of the vagrant box use `make linux`.
2. Run the vagrant setup procedure.
    ```
    vagrant up
    vagrant ssh
    ```

3. Import steam to vagrant and create the database
    ```
    ln -s /steam/steam--linux-amd64/ steam
    cd steam--linux-amd64/var/master/scripts
    sudo -u postgres createuser steam
    sudo -u postgres ./create-database.sh
    ```

4. Steam can be run using the default starting commands. (The scoring service can be run using the default procedure as well.)
    ```
    cd ../../..
    ./steam serve master \
    --superuser-name=superuser \
    --superuser-password=superuser
    ```



