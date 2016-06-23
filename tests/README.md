# Testing notes

## Prerequisites

### Install postgresql

    brew update
    brew install postgres

### Install pgAdmin (optional, useful for debugging)

Download and install from here: https://www.pgadmin.org/download/macosx.php

### Configuration (optional, useful for debugging)

Edit `/usr/local/var/postgres/postgresql.conf` and turn on verbose logging:

    log_statement = 'all'

## Setup

1. Start postgresql:

        postgres -D /usr/local/var/postgres

1. Open a new terminal window

1. Get Steam:

        s3cmd get s3://steam-release/steamY-master-darwin-amd64.tar.gz
        tar xvf steamY-master-darwin-amd64.tar.gz
        cd steam-master-darwin-amd64

1. Create application user (one-time only):

        createuser -P steam

1. Create the Steam database:

        ./var/master/scripts/create-database.sh

1. Start the compilation service:

        java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war

1. Open a new terminal window

1. Start Steam

        ./steam serve master \
            --superuser-username=superuser \
            --superuser-password=superuser

## Running tests

    git clone https://github.com/h2oai/steamY.git
    cd steamY/tests
    python example.py

TODO: Add specifics


## Teardown

1. Drop the Steam database:

        cd steam-master-darwin-amd64
        ./var/master/scripts/drop-database.sh

2. `Ctrl+C` on the steam, compilation service and postgres terminal windows.

