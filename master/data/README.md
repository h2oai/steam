# Development Notes

## Install postgresql

Install using `brew`:

    brew update
    brew install postgres

## Install pgAdmin (optional)

Download and install from here: https://www.pgadmin.org/download/macosx.php

## Configuration

Edit `/usr/local/var/postgres/postgresql.conf` and turn on verbose logging:

    log_statement = 'all'     # none, ddl, mod, all

## Start postgresql

For development:

    postgres -D /usr/local/var/postgres

OR

To have launchd start postgresql at login:

    ln -sfv /usr/local/opt/postgresql/*.plist ~/Library/LaunchAgents
    launchctl load ~/Library/LaunchAgents/homebrew.mxcl.postgresql.plist

## Create application user

    createuser -P steam

## Create steam database

    psql -U su_name postgres -a -f create-database.sql

(On OSX, `su_name` will be your OSX username)

## Drop steam database

    psql -U su_name postgres -a -f drop-database.sql


