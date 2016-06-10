# OSX Installation

## Install postgresql

Install using `brew`:

    brew update
    brew install postgres

## Install pgAdmin (optional)

Download and install from here: https://www.pgadmin.org/download/macosx.php

## Configuration (development only)

Edit `/usr/local/var/postgres/postgresql.conf` and turn on verbose logging:

    log_statement = 'all'     # none, ddl, mod, all

## Start postgresql

To have launchd start postgresql at login:

    ln -sfv /usr/local/opt/postgresql/*.plist ~/Library/LaunchAgents

Then to load postgresql now:

    launchctl load ~/Library/LaunchAgents/homebrew.mxcl.postgresql.plist

Or, if you don't want/need launchctl, you can just run:

    postgres -D /usr/local/var/postgres

## Create application user

    createuser -P steam

## Set up steam database

    psql -U su_name postgres -a -f init.sql

(On OSX, `su_name` will be your OSX username)

